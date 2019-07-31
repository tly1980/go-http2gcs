package tests

import (
	"testing"
	"os"

	"github.com/bradleyjkemp/cupaloy"

	hio "http2gcs/io"
)

var SRC_SMALL = "fixtures/f1.txt"
var SRC_BIG = "fixtures/all.jl"

func TestCopy(t *testing.T) {
	r, err := os.Open(SRC_SMALL)
	if (err != nil) {
		t.Errorf("error: %v", err)
	}

	w, err := os.Create(".test-tmp/TestCopy-f1.txt")
	if (err != nil) {
		t.Errorf("error: %v", err)
	}

	defer w.Close()
	defer r.Close()

	hash, err := hio.Copy(w, r)
	cupaloy.SnapshotT(t, hash)

	t.Logf("crc32c: %v", hash)
	if (err != nil) {
		t.Logf("error: %v", err)
	}

}


func TestGZipCopySimple(t *testing.T) {
	r, err := os.Open(SRC_SMALL)
	if (err != nil) {
		t.Errorf("error: %v", err)
	}

	w, err := os.Create(".test-tmp/TestGZipCopySimple-f1.gz")
	if (err != nil) {
		t.Errorf("error: %v", err)
	}

	defer w.Close()
	defer r.Close()

	hash, err := hio.GZipCopySimple(w, r)
	cupaloy.SnapshotT(t, hash)

	t.Logf("crc32c: %v", hash)
	if (err != nil) {
		t.Logf("error: %v", err)
	}

}

func TestGZipCopy(t *testing.T) {
	r, err := os.Open(SRC_BIG)
	if (err != nil) {
		t.Errorf("error: %v", err)
	}

	w, err := os.Create(".test-tmp/TestGZipCopy-all.jl.gz")
	if (err != nil) {
		t.Errorf("error: %v", err)
	}

	defer w.Close()
	defer r.Close()

	hash, hashZ, err, err1 := hio.GZipCopy(w, r)
	cupaloy.SnapshotT(t, []uint32 {hash, hashZ})

	t.Logf("hash: %v, hashZ: %v", hash, hashZ)
	t.Logf("err: %v, err1: %v", err, err1)

}
