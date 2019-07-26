package io

import (
  "io"
  // "cloud.google.com/go/storage"
  // "path/filepath"
  //"log"
  "hash/crc32"
  "sync"
  "compress/gzip"
  "fmt"
)

var CRC32_TBL = crc32.MakeTable(crc32.Castagnoli)

func GZipCopySimple(writer io.Writer, reader io.Reader) (uint32, error) {
  // reader --> multi(orig-hash, gzwriter)
  //       \-- gzip writer --> multi(gz-hash, writer)

  hash := crc32.New(CRC32_TBL)

  // multi writer no need to close
  mw := io.MultiWriter(hash, writer)

  gzw := gzip.NewWriter(mw)
  defer gzw.Close()

  fmt.Println("start...")
  _, err := io.Copy(gzw, reader)
  //wg.Wait()
  return hash.Sum32(), err
}

func GZipCopy(writer io.Writer, reader io.Reader) (uint32, uint32, error, error) {
  // reader --> multi(orig-hash, gzwriter)
  //       \-- gzip writer --> multi(gz-hash, writer)

  wg := &sync.WaitGroup {}
  hashZ := crc32.New(CRC32_TBL)
  hash := crc32.New(CRC32_TBL)
  pr, pw := io.Pipe()
  var err1 error = nil
  gzw := gzip.NewWriter(pw)

  mw := io.MultiWriter(hash, gzw)
  mwFinal := io.MultiWriter(hashZ, writer)

  fmt.Println("start...")
  wg.Add(1)
  // start the thread first
  go func () {
    _, err1 = io.Copy(mwFinal, pr)
    wg.Done()
  } ();

  _, err := io.Copy(mw, reader)
  gzw.Close()
  pw.Close()

  wg.Wait()
  return hash.Sum32(), hashZ.Sum32(), err, err1
}
