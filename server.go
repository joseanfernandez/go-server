package main

import (
    "log"
		"net/http"
		H "./http"
)

func films(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query()["key"]
	data := H.HttpRequest(key[0])
	w.Header().Set("Content-Type", "application/json")
  w.Write(data)
}

func main() {
		http.HandleFunc("/films", films)
		log.Fatal(http.ListenAndServe(":8081", nil))
}
