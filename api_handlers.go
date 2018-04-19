package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Check struct {
	ID            int       `json:"id" db:"id"`
	Domain        string    `json:"domain" db:"domain"`
	LastPerformed time.Time `json:"tested_at" db:"last_performed"`
	Status        bool      `json:"status" db:"is_up"`
}

type Incident struct {
	ID           int       `json:"id" db:"id"`
	CheckID      int       `json:"check_id" db:"check_id"`
	Description  string    `json:"description" db:"description"`
	DetectedDown time.Time `json:"downtime_started" db:"down_detection"`
	DetectedUp   time.Time `json:"downtime_finished" db:"up_detection"`
}

func apiHasInternalServerError(writer http.ResponseWriter) {
	writer.WriteHeader(500)

	jsonObject, _ := json.Marshal(map[string]string{
		"error": "an unknown error has occurred.",
	})

	writer.WriteHeader(500)
	writer.Write(jsonObject)
}

func apiWriteResponse(writer http.ResponseWriter, response interface{}) {
	jsonObject, err := json.Marshal(response)
	if err != nil {
		apiHasInternalServerError(writer)
		return
	}

	writer.WriteHeader(200)
	writer.Write(jsonObject)
}

func apiStatuses(writer http.ResponseWriter, request *http.Request) {
	apiWriteResponse(writer, getChecks())
}

func apiDownStatuses(writer http.ResponseWriter, request *http.Request) {
	apiWriteResponse(writer, getFailingChecks())
}

func apiDomainStatus(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	apiWriteResponse(writer, getDomainCheck(vars["domain"]))
}

func apiDomainIncidents(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	apiWriteResponse(writer, getDomainIncidents(vars["domain"]))
}
