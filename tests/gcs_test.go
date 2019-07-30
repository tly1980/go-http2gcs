package tests

import (
  "fmt"
  "testing"

  "github.com/bradleyjkemp/cupaloy"

  "http2gcs/consumer"
  "http2gcs/task"
)

func TestGCSTask(t *testing.T) {
  feedBack := consumer.Do(
    consumer.NewGCSDumper(),
    &task.Task{
      "https://gist.githubusercontent.com/tly1980/f1f1f9b99be233d15b955dd40320913a/raw/288943ac222307988c6e05a30e0a1820dcb63d4b/gcs_ut1.txt",
      "gs",
      "tt-dev",
      "ut/http2gcs/gcs_ut1.txt.gz",
  })

  cupaloy.Snapshot(feedBack)

  hashZHEX := fmt.Sprintf("%#x", feedBack.HashZ)
  fmt.Printf("hashZ: %v", hashZHEX)
}
