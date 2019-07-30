package tests

import (
	"testing"
	"os"
	// "sync"
	// "io"
	// "compress/gzip"
	hio "http2gcs/io"
)

func TestGZipCopySimple(t *testing.T) {
	r, err := os.Open("fixtures/f1.txt")
	if (err != nil) {
		t.Errorf("error: %v", err)
	}

	w, err := os.Create("f1.gz")
	if (err != nil) {
		t.Errorf("error: %v", err)
	}

	defer w.Close()
	defer r.Close()

	hash, err := hio.GZipCopySimple(w, r)

	t.Logf("crc32c: %v", hash)
	if (err != nil) {
		t.Logf("error: %v", err)
	}

}

func TestGZipCopy(t *testing.T) {
	r, err := os.Open("fixtures/all.jl")
	if (err != nil) {
		t.Errorf("error: %v", err)
	}

	w, err := os.Create("all.jl.gz")
	if (err != nil) {
		t.Errorf("error: %v", err)
	}

	defer w.Close()
	defer r.Close()

	hash, hashZ, err, err1 := hio.GZipCopy(w, r)

	t.Logf("hash: %v, hashZ: %v", hash, hashZ)
	t.Logf("err: %v, err1: %v", err, err1)

}
