package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/byuoitav/scurvy/adcp"
	"github.com/spf13/pflag"
)

// Jitter represents the number of seconds of random jitter allowed
const _jitter = 30

// MinimumWait represents the minimum number of seconds between commands
const _minimumWait = 300

func main() {

	var address string

	pflag.StringVarP(&address, "address", "a", "", "The address of the projector to test")

	pflag.Parse()

	if len(address) == 0 {
		log.Printf("No address given for a projector to test")
		os.Exit(1)
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	inputs := []string{
		"hdmi1",
		"hdmi2",
		"hdbaset1",
	}

	d := adcp.New(address, adcp.WithInputs(inputs...))

	log.Printf("Starting to dispatch events to %s...", address)

	for {
		cmd, err := d.RandomDispatch()
		if err != nil {
			log.Printf("Dispatch error!: %s", err)
		} else {
			log.Printf("Dispatch successful! Command: %s", cmd)
		}

		sleepDelay := _minimumWait + rand.Intn(_jitter)

		time.Sleep(time.Duration(sleepDelay) * time.Second)
	}
}
