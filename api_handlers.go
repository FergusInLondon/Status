package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Check struct {
	ID            int       `json:"id"`
	Domain        string    `json:"domain"`
	LastPerformed time.Time `json:"tested_at"`
	Status        bool      `json:"status"`
}

type Incident struct {
	ID           int       `json:"id"`
	CheckID      int       `json:"check_id"`
	Description  string    `json:"description"`
	DetectedDown time.Time `json:"downtime_started"`
	DetectedUp   time.Time `json:"downtime_finished"`
}

func apiStatuses(writer http.ResponseWriter, request *http.Request) {
	jsonObject, err := json.Marshal(getChecks())

	if err == nil {
		writer.WriteHeader(200)
		writer.Write(jsonObject)
		return
	}

	writer.WriteHeader(500)
}

func apiDownStatuses(writer http.ResponseWriter, request *http.Request) {
	jsonObject, err := json.Marshal(getFailingChecks())

	if err == nil {
		writer.WriteHeader(200)
		writer.Write(jsonObject)
		return
	}

	writer.WriteHeader(500)
}

func apiDomainStatus(writer http.ResponseWriter, request *http.Request) {
	jsonObject, err := json.Marshal(getDomainCheck("fergus.london"))

	if err == nil {
		writer.WriteHeader(200)
		writer.Write(jsonObject)
		return
	}

	writer.WriteHeader(500)
}

func apiDomainIncidents(writer http.ResponseWriter, request *http.Request) {
	jsonObject, err := json.Marshal(getDomainIncidents("fergus.london"))

	if err == nil {
		writer.WriteHeader(200)
		writer.Write(jsonObject)
		return
	}

	writer.WriteHeader(500)
}
