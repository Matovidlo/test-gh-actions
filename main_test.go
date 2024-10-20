package main

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	output := captureOutput(t, main)
	assert.Equal(t, "Hello world: \nHere is simple text as file\n", output)
}

func captureOutput(t *testing.T, f func()) string {
	t.Helper()
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
