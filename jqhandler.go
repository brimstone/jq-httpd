package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/brimstone/jq-httpd/jq"
)

func JqHandler(w http.ResponseWriter, clientRequest *http.Request) {
	// Report our source location
	w.Header().Add("X-Source", "https://github.com/brimstone/jq-httpd")
	// Report our return time
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// figure out if we have the right number of parameters
	parts := strings.Split(clientRequest.URL.Path, "/")
	if len(parts) != 5 {
		errorjson(w, 404, "Expected url in format /jq/urlencode(jq filter)/to/urlencode(path)")
		return
	}
	// read in all of our body
	userjson, err := ioutil.ReadAll(clientRequest.Body)
	if err != nil {
		errorjson(w, 500, "Can't read body")
		return
	}
	// actually perform the transformation
	results, err := jq.Process(userjson, parts[2])
	if err != nil {
		errorjson(w, 500, err.Error())
		return
	}

	if len(results) == 0 {
		errorjson(w, 400, "No Results")
		return
	}

	// If the /to/ bit is empty, return the transformation to the user
	if parts[4] == "" {
		// return the results to the user, for now
		w.Write(results[0])
		return
	}

	// Since /to/ is set, relay the request
	client := &http.Client{}
	serverRequest, err := http.NewRequest(clientRequest.Method, parts[4], nil)
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
	serverRequest.Header.Set("X-FORWARDED-FOR", clientRequest.RemoteAddr)

	// Actually perform the request
	// v2 will do something with the server response, check for !200 maybe
	_, err = client.Do(serverRequest)
	if err != nil {
		errorjson(w, 400, err.Error())
		return
	}

	w.Write([]byte("Successful send"))

}
