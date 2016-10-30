package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/brimstone/jq-httpd/jq"
)

func JqHandler(w http.ResponseWriter, clientRequest *http.Request, jqPattern string, serverURL string) {
	fmt.Println("ServerURL:", serverURL)
	// Report our source location
	w.Header().Add("X-Source", "https://github.com/brimstone/jq-httpd")
	// Report our return time
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// figure out if we have the right number of parameters
	// read in all of our body
	userjson, err := ioutil.ReadAll(clientRequest.Body)
	if err != nil {
		errorjson(w, 500, "Can't read body")
		return
	}
	// actually perform the transformation
	results, err := jq.Process(userjson, jqPattern)
	if err != nil {
		errorjson(w, 500, err.Error())
		return
	}

	if len(results) == 0 {
		errorjson(w, 400, "No Results")
		return
	}

	// Fake encode the result into an array, if it needs it
	result := []byte("[\n")
	if len(results) == 1 {
		result = results[0]
	} else {
		result = bytes.Join(results, []byte(",\n"))
		result = append(result, []byte("]\n")...)
	}

	// If the /to/ bit is empty, return the transformation to the user
	if serverURL == "" {
		// return the results to the user, for now
		w.Write(result)
		return
	}

	// Since /to/ is set, relay the request
	client := &http.Client{
		Timeout: time.Duration(time.Second),
	}
	serverRequest, err := http.NewRequest(
		clientRequest.Method,
		serverURL,
		bytes.NewBuffer(result),
	)
	if err != nil {
		errorjson(w, 400, err.Error())
		return
	}
	// Set all headers same as the client
	for k, vs := range clientRequest.Header {
		if k == "Content-Length" {
			continue
		}
		for _, v := range vs {
			serverRequest.Header.Set(k, v)
		}
	}
	serverRequest.Header.Set("X-Forwarded-For", clientRequest.RemoteAddr)

	// TODO use a context or something

	// Actually perform the request
	// TODO next version will do something with the server response, check for !200 maybe
	_, err = client.Do(serverRequest)
	if err != nil {
		errorjson(w, 400, err.Error())
		return
	}

	w.Write([]byte("Successful send"))

}
