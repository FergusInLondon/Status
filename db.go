package main

import (
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

var connection *sqlx.DB

func databaseConnection() (err error) {
	config := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		DBName: os.Getenv("MYSQL_DATABASE"),
		Net:    "tcp",
		Addr:   os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_POSRT"),
		Params: map[string]string{
			"parseTime": "true",
		},
	}

	connection, err = sqlx.Open("mysql", config.FormatDSN())
	return err
}

func getChecks() []Check {
	checks := []Check{}
	err := connection.Select(&checks, "SELECT * FROM checks")
	if err != nil {
		panic(err)
	}

	log.Println(checks)

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
	_, err := connection.NamedExec(`UPDATE checks SET last_performed=:time, is_up=:status WHERE id = :id`,
		map[string]interface{}{
			"time":   current.LastPerformed,
			"status": current.Status,
			"id":     current.ID,
		})

	if err != nil {
		panic(err)
	}

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
	log.Println("Downtime incident detected on ", domain.Domain)

	_, err := connection.NamedExec(`INSERT INTO incidents (check_id, down_detection) VALUES (:id, :time)`,
		map[string]interface{}{
			"id":   domain.ID,
			"time": domain.LastPerformed,
		})

	if err != nil {
		panic(err)
	}
}

func resolveIncident(domain Check) {
	log.Println("Downtime incident over on ", domain.Domain)

	incident := &Incident{}
	connection.Get(incident, `SELECT * FROM incidents 
		WHERE check_id = $1 AND up_detection = NULL ORDER BY id DESC LIMIT 1`, domain.ID)

	_, err := connection.NamedExec(`UPDATE incidents SET up_detection=:time WHERE id = :id`,
		map[string]interface{}{
			"time": domain.LastPerformed,
			"id":   incident.ID,
		})

	if err != nil {
		panic(err)
	}
}
