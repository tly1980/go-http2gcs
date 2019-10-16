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
	// add some buffer to the channel, to make it async
	ch := make(chan *task.Task, 20)
	file, _ := os.Open("fixtures/f1.txt")
	defer file.Close()

	producer.FileListToDirProducer(file, "gs://ut/a/b/c", ch)
	var allTasks []*task.Task
	close(ch)
	for theTask := range ch {
		allTasks = append(allTasks, theTask)
	}

	cupaloy.SnapshotT(t, allTasks)
}

/*
func TestFileListToDirProducer(t *testing.T) {
	ch := make(chan *task.Task)
	file, _ := os.Open("fixtures/task1.csv")
	defer file.Close()

	go producer.FileListToDirProducer(file, "gs://ut/a/b/c", ch)
	l1, l2, l3, l4 := <-ch, <-ch, <-ch, <-ch
	allTasks := []*task.Task {l1, l2, l3, l4}
	cupaloy.SnapshotT(t, allTasks)
}
*/

func TestCSVFileToFileProducer(t *testing.T) {
	// add some buffer to the channel, to make it async
	ch := make(chan *task.Task, 20)
	file, _ := os.Open("fixtures/task1.csv")
	defer file.Close()

	producer.CSVFileToFileProducer(file, ch)
	var allTasks []*task.Task
	close(ch)
	for theTask := range ch {
		allTasks = append(allTasks, theTask)
	}
	//l1, l2, l3 := <-ch, <-ch, <-ch
	// := {l1, l2, l3}
	cupaloy.SnapshotT(t, allTasks)
}
