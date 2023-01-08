package main

import (
	"fmt"
	"net/http"
	"log"
)

const port string = ":4000"

func main() {
	http.HandleFunc("/", index)
	log.Println("Listening on localhost ", port)
	http.ListenAndServe(port, nil)
}

func index (w http.ResponseWriter, r *http.Request){
   fmt.Fprintf(w , "Hello")
}