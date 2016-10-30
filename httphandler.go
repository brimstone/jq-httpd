package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type errorJSON struct {
	Error string `json:"error"`
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
	// Report our source location
	w.Header().Add("X-Source", "https://github.com/brimstone/jq-httpd")
	// Report our LICENSE
	w.Header().Add("X-License", "AGPLv3 http://www.gnu.org/licenses/agpl-3.0.txt")
	parts := strings.Split(request.URL.Path, "/")
	if parts[1] == "jq" {
		if len(parts) < 5 {
			errorjson(w, 404, "Expected url in format /jq/urlencode(jq filter)/to/urlencode(path)")
			return
		}
		JqHandler(w,
			request,
			parts[2],
			request.URL.Path[len(parts[2])+8:],
		)
	}
}
