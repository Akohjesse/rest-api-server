package main

import (
	"net/http"
	"log"
	// "fmt"
)

const port string = ":4000"

func main() {
	http.HandleFunc("/", index)
	log.Println("Listening on localhost ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func index (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		// fmt.Fprintf(w, "{message: '%q method requested'}", r.Method)
		w.Write([]byte(`{message: "GET method requested"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{message: "POST method requested"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))

	}
	
}