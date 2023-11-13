package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	vscontrol "github.com/byuoitav/scurvy/visionary-solutions-control"
	"github.com/spf13/pflag"
)

// Jitter represents the number of seconds of random jitter allowed
const _jitter = 5

// MinimumWait represents the minimum number of seconds between commands
const _minimumWait = 30

func main() {

	var (
		address string
		port    string
	)

	pflag.StringVarP(&address, "address", "a", "", "The address of the VSControl microservice to test")
	pflag.StringVarP(&port, "port", "p", "", "The port of the VSControl microservice")
	encoders := pflag.CommandLine.StringArrayP("encoder", "e", []string{}, "The address of an encoder to test connections")
	decoders := pflag.CommandLine.StringArrayP("decoder", "d", []string{}, "The address of an encoder to test connections")

	pflag.Parse()

	if len(address) == 0 {
		log.Printf("No address given")
		os.Exit(1)
	}

	// Seed the random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// make device
	if len(*encoders) == 0 {
		log.Printf("At least one encoder is required")
		os.Exit(1)
	}

	if len(*decoders) == 0 {
		log.Printf("At least one decoder is required")
		os.Exit(1)
	}

	log.Print(address)
	log.Print(port)
	log.Print(encoders)
	log.Print(decoders)

	vs := vscontrol.New(address, port, encoders, decoders)

	for {
		//random dispatch
		cmd, err := vs.RandomDispatch()
		if err != nil {
			log.Printf("Dispatch failed: %v", err.Error())
		} else {
			log.Printf("Dispatch successful. Command: %v", cmd)
		}

		sleepDelay := _minimumWait + r.Intn(_jitter)

		time.Sleep(time.Duration(sleepDelay) * time.Second)
	}
}
