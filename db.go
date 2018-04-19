package main

import "time"

var (
	exampleChecks = []Check{
		Check{ID: 0, Domain: "example.com", LastPerformed: time.Now(), Status: true},
		Check{ID: 1, Domain: "fergus.london", LastPerformed: time.Now(), Status: false},
		Check{ID: 2, Domain: "google.com", LastPerformed: time.Now(), Status: true},
		Check{ID: 3, Domain: "facebook.com", LastPerformed: time.Now(), Status: true},
		Check{ID: 4, Domain: "github.com", LastPerformed: time.Now(), Status: false},
	}
	exampleIncidents = []Incident{
		Incident{ID: 0, CheckID: 1, Description: "It went down.", DetectedDown: time.Now(), DetectedUp: time.Now()},
		Incident{ID: 1, CheckID: 4, Description: "It went down.", DetectedDown: time.Now(), DetectedUp: time.Now()},
		Incident{ID: 2, CheckID: 1, Description: "It went down.", DetectedDown: time.Now(), DetectedUp: time.Now()},
		Incident{ID: 3, CheckID: 4, Description: "It went down.", DetectedDown: time.Now(), DetectedUp: time.Now()},
	}
)

func getChecks() []Check {
	return exampleChecks
}

func getFailingChecks() []Check {
	failingChecks := []Check{}
	for _, check := range exampleChecks {
		if !check.Status {
			failingChecks = append(failingChecks, check)
		}
	}

	return failingChecks
}

func getDomainCheck(domain string) Check {
	var check Check

	for _, check := range exampleChecks {
		if check.Domain == domain {
			return check
		}
	}

	return check
}

func getDomainIncidents(domain string) []Incident {
	domainIncidents := []Incident{}

	check := getDomainCheck(domain)

	for _, incident := range exampleIncidents {
		if incident.CheckID == check.ID {
			domainIncidents = append(domainIncidents, incident)
		}
	}

	return domainIncidents
}
