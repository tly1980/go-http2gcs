package consumer

import (
  "bufio"
  "net/http"
  "time"
  "context"
  // third parties

  "cloud.google.com/go/storage"
  "github.com/cenkalti/backoff"
  cache "github.com/patrickmn/go-cache"
  log "github.com/sirupsen/logrus"

  hio "http2gcs/io"
  "http2gcs/task"
)

var LOG = log.New()

const EXPIRATION = 10 * time.Minute

type Dumper interface {
  GetWriter(task *task.Task) (*storage.Writer, error)
}


type GCSDumper struct {
  bktCache *cache.Cache
}

func NewGCSDumper() *GCSDumper {
  bktCache := cache.New(EXPIRATION, 2 * EXPIRATION)
  return &GCSDumper{ bktCache }
}

func (self *GCSDumper) GetBucket(bktName string) (*storage.BucketHandle, error) {

  item, found := self.bktCache.Get(bktName)
  var bkt *storage.BucketHandle
  if found {
    bkt = item.(*storage.BucketHandle)
    return bkt, nil
  }

  ctx := context.Background()
  client, err := storage.NewClient(ctx)
  if err != nil {
    LOG.Error("Failed to create GCS storage.Client: %v", err)
    return nil, err
  }
  bkt = client.Bucket(bktName)
  self.bktCache.Add(bktName, bkt, EXPIRATION)
  return bkt, nil
}

func (self * GCSDumper) GetWriter(task *task.Task) (*storage.Writer, error) {
  bkt, err := self.GetBucket(task.DestBkt)
  if (err != nil) {
    return nil, err
  }
  ctx := context.Background()
  obj := bkt.Object(task.DestKey)
  return obj.NewWriter(ctx), nil
}

func Do(dumper *GCSDumper, theTask *task.Task) (*task.FeedBack) {
  LOG.Error("Do: task=", theTask)
  var resp *http.Response
  var writer *storage.Writer
  feedBack := &task.FeedBack {theTask, 0, 0, nil}

  expBackoff := backoff.NewExponentialBackOff()
  expBackoff.MaxElapsedTime = time.Duration(time.Minute *5 )

  expBackoff.Reset()
  err := backoff.Retry(func () error {
    var err0 error
    writer, err0 = dumper.GetWriter(theTask)
    return err0
  }, expBackoff)
  if err != nil {
    LOG.Error("Stage1: task=", theTask, "err=", err)
    feedBack.Err = err
    return feedBack
  }
  defer writer.Close()

  expBackoff.Reset()
  err = backoff.Retry(func() error {
    var err0 error
    resp, err0 = http.Get(theTask.Src)
    return err0
  }, expBackoff)
  if err != nil {
    LOG.Error("Stage1: task=", theTask, "err=", err)
    feedBack.Err = err
    return feedBack
  }
  defer resp.Body.Close()

  expBackoff.Reset()
  err = backoff.Retry(func() error {
    bWriter := bufio.NewWriter(writer)
    bReader := bufio.NewReader(resp.Body)
    hash, hashZ, err0, err1 := hio.GZipCopy(bWriter, bReader)
    feedBack.Hash = hash
    feedBack.HashZ = hashZ
    bWriter.Flush()
    if (err0 != nil) {
      return err0
    }
    if (err1 != nil) {
      return err1
    }
    return nil
  }, expBackoff)
  if err != nil {
    feedBack.Err = err
    LOG.Error("Stage3 - Got err: %v, %v", theTask, err)
  }

  return feedBack
}


func StartProc(workerCount int, tasks chan *task.Task, feedBacks chan *task.FeedBack) {
  for i := 0; i < workerCount; i++ {
    go func() {
      dump := NewGCSDumper()
      for t := range tasks {
        feedBacks <- Do(dump, t)
      }
    } ()
  }
}

