package jq_test

import (
	"bytes"
	"testing"

	"github.com/brimstone/jq-httpd/jq"
)

func TestProcess(t *testing.T) {
	results, err := jq.Process([]byte("{\"name\": \"pickles\"}"), ".name")
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 {
		t.Fatal("Expected exactly one result")
	}
	if bytes.Compare(results[0], []byte("\"pickles\"")) != 0 {
		t.Fatal("Expected", []byte("\"pickles\""), "and got", results[0])
	}
}

func TestBad(t *testing.T) {
	_, err := jq.Process([]byte("{"), ".")
	if err == nil {
		t.Fatal("Expected parse error")
	}
}
