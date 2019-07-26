package producer

import (
  "bufio"
  "log"
  "os"
)

func FileProducer(fPath string, task chan string) {
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
