package main

import (
	"log"
	"time"
)

var (
	duration time.Duration
	checks   []Check
)

func perform_checks() (err error) {
	duration, err = time.ParseDuration("15s")
	if err != nil {
		return err
	}

	log.Println("Starting monitoring routine..")

	go func() {
		for {
			check()
			time.Sleep(duration)
		}
	}()

	log.Println("Done.")

	return err
}

func check() {
	log.Println("Performing check.")
}
