package tests

import (
	"testing"
	// "github.com/stretchr/testify/assert"
	// "http2gcs/task"
	// "http2gcs/producer"
)

func TestProducer(t *testing.T) {
	t.Logf("aaa\n")
}

/*
func TestProducer(t *testing.T) {
	ch := make(chan *task.Task)
	go producer.FileProducer("fixtures/f1.txt", "gs://ut/a/b/c", ch)
	l1, l2, l3 := <-ch, <-ch, <-ch
	assert.Equal(t, l1, "http://a.com/", "The two words should be the same.")
	assert.Equal(t, l2, "http://bb.com/", "The two words should be the same.")
	assert.Equal(t, l3, "http://ccc.com/", "The two words should be the same.")

} */
