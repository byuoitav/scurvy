package atlona5x2

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	atuhdsw52ed "github.com/byuoitav/atlona/AT-UHD-SW-52ED"
	"github.com/byuoitav/scurvy"
)

var _MUTES = []string{
	"HDMI",
	"HDBT",
	"ANALOG",
}

var _INPUTS = []string{
	"1",
	"2",
	"3",
	"4",
	"5",
}

type Atlona5x2 struct {
	logger scurvy.Logger
	device *atuhdsw52ed.AtlonaVideoSwitcher5x1
}

type NullLogger struct{}

func (l *NullLogger) Debugf(format string, a ...interface{}) {}
func (l *NullLogger) Infof(format string, a ...interface{})  {}
func (l *NullLogger) Warnf(format string, a ...interface{})  {}
func (l *NullLogger) Errorf(format string, a ...interface{}) {}

func New(username, password, address string, opts ...Option) scurvy.RandomDispatcher {
	d := Atlona5x2{}

	// Create device
	sw := atuhdsw52ed.AtlonaVideoSwitcher5x1{
		Username: username,
		Password: password,
		Address:  address,
		Logger:   &NullLogger{},
	}

	d.device = &sw

	// Apply Options
	for _, opt := range opts {
		opt(&d)
	}

	rand.Seed(time.Now().UnixNano())

	return &d
}

func (d *Atlona5x2) RandomDispatch() (string, error) {
	var err error = nil
	var cmd string

	// Random dispatch
	switch rand.Intn(5) {
	case 0: // Get Volume
		cmd = "Get Volume"
		_, err = d.device.Volumes(context.TODO(), []string{"block"})

	case 1: // Set Volume
		vol := rand.Intn(101)
		cmd = fmt.Sprintf("Set Volume (%d)", vol)
		err = d.device.SetVolume(context.TODO(), "block", vol)

	case 2: // Get Input
		cmd = "Get Input"
		_, err = d.device.AudioVideoInputs(context.TODO())

	case 3: // Set Input
		input := _INPUTS[rand.Intn(len(_INPUTS))]
		cmd = fmt.Sprintf("Set Input (%s)", input)
		err = d.device.SetAudioVideoInput(context.TODO(), "OUT", input)

	default: // Skip
		cmd = "Skip"
	}

	// If an error occurred
	if err != nil {
		return cmd, fmt.Errorf("Error while trying to issue %s command: %w", cmd, err)
	}

	return cmd, nil
}
