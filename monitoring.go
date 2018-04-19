package main

import (
	"log"
	"net"
	"time"

	fastping "github.com/tatsushid/go-fastping"
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
		checks = getChecks()
		for {
			check()
			time.Sleep(duration)
		}
	}()

	log.Println("Done.")

	return err
}

func check() {
	log.Println("Performing checks...")

	for idx, check := range checks {
		log.Println("Processing check for... ", check.Domain)

		ping := fastping.NewPinger()
		addr, err := net.ResolveIPAddr("ip4:icmp", check.Domain)
		if err != nil {
			panic(err)
		}

		ping.AddIPAddr(addr)

		// We've recieved a response, therefore our host is up.
		recievedResponse := false
		ping.OnRecv = func(addr *net.IPAddr, t time.Duration) {
			recievedResponse = true
		}

		// This is really "onComplete".
		ping.OnIdle = func() {
			newCheck := check
			newCheck.Status = recievedResponse
			newCheck.LastPerformed = time.Now()

			updateDomain(check, newCheck)
			checks[idx] = newCheck
		}

		err = ping.Run()
		if err != nil {
			panic(err)
		}
	}
}
