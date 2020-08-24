package epson

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/byuoitav/epson-driver"
	"github.com/byuoitav/scurvy"
)

type RandomDispatcher struct {
	inputs                []string
	logger                scurvy.Logger
	poweredOn             bool
	device                *epson.Projector
	lastSuccessfulCommand time.Time
}

func New(address string, opts ...Option) *RandomDispatcher {
	d := RandomDispatcher{
		inputs:                []string{},
		logger:                &scurvy.NullLogger{},
		poweredOn:             false,
		lastSuccessfulCommand: time.Time{},
	}

	// Apply Options
	for _, opt := range opts {
		opt(&d)
	}

	rand.Seed(time.Now().UnixNano())

	d.device = epson.NewProjector(address)

	return &d
}

func (d *RandomDispatcher) RandomDispatch() (string, error) {

	// Always power on if the device is off
	if !d.poweredOn {
		cmd := "Set Power (on)"
		err := d.device.SetPower(context.TODO(), "on")
		if err != nil {
			return cmd, fmt.Errorf("Error while trying to issue Set Power (on) command: %w", err)
		}
		d.poweredOn = true

		return cmd, nil
	}

	var err error = nil
	var cmd string
	// Random dispatch
	switch rand.Intn(9) {
	case 0: // Get Volume
		cmd = "Get Volume"
		_, err = d.device.GetVolumeByBlock(context.TODO(), "block")

	case 1: // Set Volume
		vol := rand.Intn(101)
		cmd = fmt.Sprintf("Set Volume (%d)", vol)
		err = d.device.SetVolumeByBlock(context.TODO(), "block", vol)

	case 2: // Get Input
		cmd = "Get Input"
		_, err = d.device.GetInput(context.TODO())

	case 3: // Set Input
		if len(d.inputs) > 0 {
			input := d.inputs[rand.Intn(len(d.inputs))]
			cmd = fmt.Sprintf("Set Input (%s)", input)
			err = d.device.SetInput(context.TODO(), input)
			break
		}
		fallthrough // if we didn't happen to give any inputs, do the next case

	case 4: // Get Blanked
		cmd = "Get Blanked"
		_, err = d.device.GetBlanked(context.TODO())

	case 5: // Set Blanked
		blanked := false
		if rand.Intn(2) == 1 {
			blanked = true
		}
		cmd = fmt.Sprintf("Set Blanked (%t)", blanked)
		err = d.device.SetBlanked(context.TODO(), blanked)

	case 6: // Get Power
		cmd = "Get Power"
		_, err = d.device.GetPower(context.TODO())

	case 7: // Set Power
		power := "on"
		if rand.Intn(2) == 1 {
			power = "off"
		}
		cmd = fmt.Sprintf("Set Power (%s)", power)
		err = d.device.SetPower(context.TODO(), power)

		// Update poweredOn accordingly
		if err == nil {
			if power == "on" {
				d.poweredOn = true
			} else {
				d.poweredOn = false
			}
		}

	default: // Skip
		cmd = "Skip"

	}

	// If an error occurred
	if err != nil {
		return cmd, fmt.Errorf("Error while trying to issue %s command: %w", cmd, err)
	}

	// Log a successful event
	d.lastSuccessfulCommand = time.Now()

	return cmd, nil
}
