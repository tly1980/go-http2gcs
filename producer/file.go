package producer

import (
  "bufio"
  "net/url"
  "path"
  "strings"
  "io"

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

/*
func CSVFileToFileProducer(fPath string, tasks chan *task.Task) error {
  uParse, err := url.Parse(baseUri)

  if err != nil {
    return err
  }

  destScheme := uParse.Scheme
  destBkt := uParse.Host
  destBase := strings.Trim(uParse.Path, "/")
  file, err := os.Open(fPath)
  if err != nil {
    log.Error(err)
    return err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
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
*/
