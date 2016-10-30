package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	mymux := &httphandler{}
	fmt.Println("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, mymux))
}
