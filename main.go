package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	apachelog "github.com/lestrrat/go-apache-logformat"
)

var GitSummary string

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	mymux := &httphandler{}
	fmt.Println("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, apachelog.CombinedLog.Wrap(mymux, os.Stderr)))
}
