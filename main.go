package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/jq/", JqHandler)
	mymux := &httphandler{}
	fmt.Println("Listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", mymux))
}
