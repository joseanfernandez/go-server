package main

import (
	"fmt"
	"log"
	"net/http"
	H "./http"
)

func films(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8100")
	key := r.URL.Query()["key"]
	fmt.Println("Getting: " + key[0])
	data := H.HttpRequest(key[0])
	w.Header().Set("Content-Type", "application/json")
  w.Write(data)
}

func main() {
		http.HandleFunc("/films", films)
		fmt.Println("Runing on http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
}
