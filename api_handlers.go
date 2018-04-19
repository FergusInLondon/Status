package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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
