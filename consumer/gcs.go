import (
  "bufio"
  "log"
  "os"
  "math"
  "time"
  "net/http"
  "hash/crc32"
  // third parties
  "cloud.google.com/go/storage"
)




func runOne(maxRetry int, url string, destBase string, client *storage.Client) {
  sleepInterval := 200

  for i:=1; i < maxRetry; i++ {
    resp, err := http.Get(url)

    if err == nil {
      tempFile, err1 := ioutil.TempFile("http2gcs", "tfile")
      crc32qTable := crc32.MakeTable(crc32.Constants.Castagnoli)
      crc32c.Update
      resp.Body.Close()
      break
    } else {
      sleepMills := math.Pow(2, i) * sleepInterval
      continue
    }

  }
}

func run(task chan string, feedBack chan string) {

}

func BatchRun(workerNum int, task chan string, feedBack chan string) {
 for url := range chan {

  }
}
