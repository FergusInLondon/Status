package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var connection *sqlx.DB

func databaseConnection() (err error) {
	connection, err = sqlx.Open("mysql", "status_app:status_password@/status")
	return err
}

func getChecks() []Check {
	checks := []Check{}
	connection.Select(&checks, "SELECT * FROM checks")

	return checks
}

func getFailingChecks() []Check {
	failingChecks := []Check{}
	connection.Select(&failingChecks, "SELECT * FROM checks WHERE NOT is_up")

	return failingChecks
}

func getDomainCheck(domain string) Check {
	var check Check
	connection.Select(&check, "SELECT * FROM checks WHERE domain = '$1'", domain)

	return check
}

func getDomainIncidents(domain string) []Incident {
	domainIncidents := []Incident{}
	connection.Select(&domainIncidents, "SELECT i.* FROM incidents i LEFT JOIN checks c ON i.check_id = c.id WHERE c.domain = '$1'", domain)

	return domainIncidents
}

func updateDomain(previous, current Check) {
	// Update the domain

	var incidentHasBegan = previous.Status && !current.Status
	if incidentHasBegan {
		createIncident(current)
	}

	var incidentHasFinished = !previous.Status && current.Status
	if incidentHasFinished {
		resolveIncident(current)
	}
}

func createIncident(domain Check) {

}

func resolveIncident(domain Check) {

}
