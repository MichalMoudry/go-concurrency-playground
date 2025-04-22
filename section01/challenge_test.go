package main

import (
	"io"
	"os"
	"testing"
)

func Test_challange(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	challange()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if output != "Hello, universe!\nHello, cosmos!\nHello, world!\n" {
		t.Errorf("Output does not contain expected content\n")
	}
}
