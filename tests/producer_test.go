package tests

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	// "github.com/stretchr/testify/assert"
	"http2gcs/task"
	"http2gcs/producer"
)

func TestProducer(t *testing.T) {
	ch := make(chan *task.Task)
	go producer.FileProducer("fixtures/f1.txt", "gs://ut/a/b/c", ch)
	l1, l2, l3, l4 := <-ch, <-ch, <-ch, <-ch
	allTasks := []*task.Task {l1, l2, l3, l4}
	cupaloy.SnapshotT(t, allTasks)
}

