package producer

import (
  "bufio"
  "encoding/csv"
  "io"
  "net/url"
  "path"
  "strings"

  log "github.com/sirupsen/logrus"
  "http2gcs/task"
)

var LOG = log.New()

func FileListToDirProducer(reader io.Reader, baseUri string, tasks chan *task.Task) error {
  uParse, err := url.Parse(baseUri)

  if err != nil {
    return err
  }

  destScheme := uParse.Scheme
  destBkt := uParse.Host
  destBase := strings.Trim(uParse.Path, "/")

  scanner := bufio.NewScanner(reader)
  for scanner.Scan() {
    src := scanner.Text()
    fName := path.Base(src)
    t := &task.Task{
        src,
        destScheme,
        destBkt,
        path.Join(destBase, fName),
    }
    tasks <- t
  }

  if err = scanner.Err(); err != nil {
    log.Fatal(err)
    return err
  }

  return nil
}


func CSVFileToFileProducer(reader io.Reader,  tasks chan *task.Task) error {
  csvReader := csv.NewReader(reader)
  for {
    record, err := csvReader.Read()

    switch err {
      case io.EOF:
        return nil
      case nil:
        src := record[0]
        dest := record[1]
        uParse, err1 := url.Parse(dest)
        if err1 != nil {
          return err1
        }
        destScheme := uParse.Scheme
        destBkt := uParse.Host
        destBase := strings.TrimLeft(uParse.Path, "/")
        var destKey string
        if (strings.HasSuffix(destBase, "/")) {
          destKey = path.Join(destBase, path.Base(src))
        } else {
          destKey = destBase
        }

        tasks <- &task.Task{ src, destScheme, destBkt, destKey}
      default:
        return err
    }
  }
  return nil
}
