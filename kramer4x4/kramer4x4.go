package kramer4x4

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/byuoitav/kramer/protocol3000"
	"github.com/byuoitav/scurvy"
)

var _INPUTS = []string{
	"1",
	"2",
	"3",
	"4",
}

var _OUTPUTS = []string{
	"1",
	"2",
	"3",
	"4",
}

type Kramer4x4 struct {
	logger scurvy.Logger
	device *protocol3000.Device
}

type NullLogger struct{}

func (l *NullLogger) Debugf(format string, a ...interface{}) {}
func (l *NullLogger) Infof(format string, a ...interface{})  {}
func (l *NullLogger) Warnf(format string, a ...interface{})  {}
func (l *NullLogger) Errorf(format string, a ...interface{}) {}

func New(address string, opts ...Option) scurvy.RandomDispatcher {
	d := Kramer4x4{}

	// Create device
	d.device = protocol3000.New(address)

	// Apply Options
	for _, opt := range opts {
		opt(&d)
	}

	rand.Seed(time.Now().UnixNano())

	return &d
}

func (d *Kramer4x4) RandomDispatch() (string, error) {
	var err error = nil
	var cmd string

	// Random dispatch
	switch rand.Intn(4) {
	case 0: // Get Input
		cmd = "Get Input"
		_, err = d.device.AudioVideoInputs(context.TODO())

	case 1: // Set Input
		input := _INPUTS[rand.Intn(len(_INPUTS))]
		output := _OUTPUTS[rand.Intn(len(_OUTPUTS))]
		cmd = fmt.Sprintf("Set Input (%s) -> (%s)", input, output)
		err = d.device.SetAudioVideoInput(context.TODO(), output, input)

	case 2: // Device Health
		cmd = "Check device health"
		err = d.device.Healthy(context.TODO())

	default: // Skip
		cmd = "Skip"
	}

	// If an error occurred
	if err != nil {
		return cmd, fmt.Errorf("error while trying to issue %s command: %w", cmd, err)
	}

	return cmd, nil
}
