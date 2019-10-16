package main

import (
  "bufio"
  "flag"
  "log"
  "os"
  // "net/http"
)

var src  = flag.String("src", "", "A file contains URLs")
var dest = flag.String("dest", "", "A gcs base url for destination")
var httpConfig = flag.String("httpConfig", "A yaml file for http config")

func requestMaker() {

}


func Producer(fPath string, task chan string) {
  file, err := os.Open(fPath)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    task <- scanner.Text()
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}

func main() {

}
