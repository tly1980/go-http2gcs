package tests

import (
	"os"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	// "github.com/stretchr/testify/assert"
	"http2gcs/task"
	"http2gcs/producer"
)

func TestFileListToDirProducer(t *testing.T) {
	ch := make(chan *task.Task)
	file, _ := os.Open("fixtures/f1.txt")
	defer file.Close()

	go producer.FileListToDirProducer(file, "gs://ut/a/b/c", ch)
	l1, l2, l3, l4 := <-ch, <-ch, <-ch, <-ch
	allTasks := []*task.Task {l1, l2, l3, l4}
	cupaloy.SnapshotT(t, allTasks)
}

