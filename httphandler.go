package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorJSON struct {
	Error string `json:'error'`
}

type httphandler struct {
}

func errorjson(w http.ResponseWriter, code int, errString string) {
	w.WriteHeader(code)
	ej := &errorJSON{
		Error: errString,
	}
	out, err := json.Marshal(ej)
	if err != nil {
		w.Write([]byte("How did this happen?"))
		return
	}
	w.Write(out)
}

func (h httphandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path)
}
