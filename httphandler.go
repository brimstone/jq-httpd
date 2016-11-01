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

	// handle version regardless of… version
	if parts[1] == "version" {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("{" +
			"\"version\": \"" + GitSummary + "\"," +
			"\"api_version\": \"v1\"" +
			"}"))
		return
	}

	// only handle v1 corrently
	if parts[1] != "v1" {
		w.WriteHeader(400)
		w.Write([]byte("Unsupported Endpoint"))
		return
	}

	// only handle /jq/
	if parts[2] == "jq" {
		if len(parts) < 6 {
			errorjson(w, 404, "Expected url in format /jq/urlencode(jq filter)/[to|jq]/urlencode(path)")
			return
		}
		JqHandler(w,
			request,
			parts[3],
			parts[4],
			request.URL.Path[len(parts[3])+11:],
		)
		return
	}
}
