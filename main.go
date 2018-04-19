package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/status", get_statuses).Methods("GET")
	router.HandleFunc("/status/down", get_down_statuses).Methods("GET")
	router.HandleFunc("/status/service/{domain:[a-z]+}", get_domain_statuses).Methods("GET")
	router.HandleFunc("/status/service/{domain:[a-z]+}/incident/{id:[0-9]+}", get_incident).Methods("GET")

	http.Handle("/", router)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
