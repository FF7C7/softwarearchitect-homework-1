package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health/", HealthCheckHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}


func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "OK"}`)
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
}
