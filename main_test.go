package main_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	main "github.com/brimstone/jq-httpd"
)

func TestJqHandler(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://example.com/jq/./to/asdf",
		bytes.NewBufferString("{\"name\":\"pickles\"}"),
	)
	w := httptest.NewRecorder()
	main.JqHandler(w, req)
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", string(body))
}
