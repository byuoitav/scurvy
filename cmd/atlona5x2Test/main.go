package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/byuoitav/scurvy/atlona5x2"
	"github.com/spf13/pflag"
)

// Jitter represents the number of seconds of random jitter allowed
const _jitter = 5

// MinimumWait represents the minimum number of seconds between commands
const _minimumWait = 30

func main() {

	var (
		username string
		password string
		address  string
	)

	pflag.StringVarP(&username, "username", "u", "", "The username for the atlona 5x2")
	pflag.StringVarP(&password, "password", "p", "", "The password for the atlona 5x2")
	pflag.StringVarP(&address, "address", "a", "", "The address of the switcher to test")

	pflag.Parse()

	if len(address) == 0 {
		log.Printf("No address given for a switcher to test")
		os.Exit(1)
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	d := atlona5x2.New(username, password, address)

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
