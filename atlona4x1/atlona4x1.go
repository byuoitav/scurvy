package atlona4x1

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	at "github.com/byuoitav/atlona/AT-JUNO-451-HDBT"
	"github.com/byuoitav/scurvy"
)

var _INPUTS = []string{
	"0",
	"1",
	"2",
	"3",
}

type Atlona4x1 struct {
	logger scurvy.Logger
	device *at.AtlonaVideoSwitcher4x1
}

type NullLogger struct{}

func (l *NullLogger) Debugf(format string, a ...interface{}) {}
func (l *NullLogger) Infof(format string, a ...interface{})  {}
func (l *NullLogger) Warnf(format string, a ...interface{})  {}
func (l *NullLogger) Errorf(format string, a ...interface{}) {}

func New(username, password, address string, opts ...Option) scurvy.RandomDispatcher {
	d := Atlona4x1{}

	// Create device
	sw := at.AtlonaVideoSwitcher4x1{
		Username: username,
		Password: password,
		Address:  address,
	}

	d.device = &sw

	// Apply Options
	for _, opt := range opts {
		opt(&d)
	}

	rand.Seed(time.Now().UnixNano())

	return &d
}

func (d *Atlona4x1) RandomDispatch() (string, error) {
	var err error = nil
	var cmd string

	// Random dispatch
	switch rand.Intn(4) {
	case 0: // Get Input
		cmd = "Get Input"
		_, err = d.device.AudioVideoInputs(context.TODO())

	case 1: // Set Input
		input := _INPUTS[rand.Intn(len(_INPUTS))]
		cmd = fmt.Sprintf("Set Input (%s)", input)
		err = d.device.SetAudioVideoInput(context.TODO(), "0", input)

	case 2: // Device Info
		cmd = "Get Device info"
		_, err = d.device.Info(context.TODO())

	default: // Skip
		cmd = "Skip"
	}

	// If an error occurred
	if err != nil {
		return cmd, fmt.Errorf("Error while trying to issue %s command: %w", cmd, err)
	}

	return cmd, nil
}
