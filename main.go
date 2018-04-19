package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/status", apiStatuses).Methods("GET")
	router.HandleFunc("/status/down", apiDownStatuses).Methods("GET")
	router.HandleFunc("/status/service/{domain:[a-z]+}", apiDomainStatus).Methods("GET")
	router.HandleFunc("/status/service/{domain:[a-z]+}/incidents", apiDomainIncidents).Methods("GET")

	http.Handle("/", router)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
