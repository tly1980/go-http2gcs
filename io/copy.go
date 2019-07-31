package io

import (
  "io"
  // "cloud.google.com/go/storage"
  // "path/filepath"
  //"log"
  log "github.com/sirupsen/logrus"
  "hash/crc32"
  "sync"
  "compress/gzip"
  // "fmt"
)

var LOG = log.New()

var CRC32_TBL = crc32.MakeTable(crc32.Castagnoli)

func CopySimple(writer io.Writer, reader io.Reader) (uint32, error) {
  hash := crc32.New(CRC32_TBL)
  mw := io.MultiWriter(hash, writer)
  LOG.Info("start...")
  n, err := io.Copy(mw, reader)
  LOG.Info("bytes copied", n)
  //wg.Wait()
  return hash.Sum32(), err
}

func GZipCopySimple(writer io.Writer, reader io.Reader) (uint32, error) {
  hash := crc32.New(CRC32_TBL)
  // buf := bytes.Buffer {}
  // multi writer no need to close
  gzw := gzip.NewWriter(writer)
  mw := io.MultiWriter(hash, gzw)

  defer gzw.Close()

  LOG.Info("start...")
  n, err := io.Copy(mw, reader)
  LOG.Info("bytes copied", n)
  //wg.Wait()
  return hash.Sum32(), err
}

func GZipCopy(writer io.Writer, reader io.Reader) (uint32, uint32, error, error) {
  // reader --> multi(orig-hash, gzwriter)
  //       \-- gzip writer --> multi(gz-hash, writer)

  var err1 error = nil
  // buf := bytes.Buffer {}
  wg := &sync.WaitGroup {}
  hash := crc32.New(CRC32_TBL)
  hashZ := crc32.New(CRC32_TBL)
  pr, pw := io.Pipe()

  gzw := gzip.NewWriter(pw)

  mw := io.MultiWriter(hash, gzw)
  mwFinal := io.MultiWriter(hashZ, writer)

  LOG.Info("start...")
  wg.Add(1)

  // start the thread first
  go func () {
    var n int64
    n, err1 = io.Copy(mwFinal, pr)
    LOG.Info("n: ", n)
    wg.Done()
  } ();

  n0, err := io.Copy(mw, reader)
  LOG.Info("n0: ", n0)
  gzw.Close()
  pw.Close()

  wg.Wait()
  return hash.Sum32(), hashZ.Sum32(), err, err1
}
