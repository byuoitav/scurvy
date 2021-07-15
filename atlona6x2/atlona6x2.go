package atlona6x2

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	atomeps62 "github.com/byuoitav/atlona/AT-OME-PS62"
	"github.com/byuoitav/scurvy"
)

var _BLOCKS = []string{
	"zoneOut1Analog",
	"zoneOut2Analog",
	"zoneOut1Digital",
	"zoneOut2Digital",
}

var _MUTES = []string{
	"zoneOut1Analog",
	"zoneOut2Analog",
	"zoneOut1Digital",
	"zoneOut2Digital",
}

var _INPUTS = []string{
	"1",
	"2",
	"3",
	"4",
	"5",
}

var _OUTPUTS = []string{
	"hdmiOutA",
	"hdmiOutB",
}

type Atlona6x2 struct {
	logger scurvy.Logger
	device *atomeps62.AtlonaVideoSwitcher6x2
}

type NullLogger struct{}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (l *NullLogger) Debugf(format string, a ...interface{}) {}
func (l *NullLogger) Infof(format string, a ...interface{})  {}
func (l *NullLogger) Warnf(format string, a ...interface{})  {}
func (l *NullLogger) Errorf(format string, a ...interface{}) {}

func boo() bool {
	return rand.Intn(2) == 0
}

func New(username, password, address string, opts ...Option) scurvy.RandomDispatcher {
	d := Atlona6x2{}

	// Create device
	sw := atomeps62.AtlonaVideoSwitcher6x2{
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

func (d *Atlona6x2) RandomDispatch() (string, error) {
	var err error = nil
	var cmd string

	// Random dispatch
	switch rand.Intn(7) {
	case 0: // Get Volume
		cmd = "Get Volume"
		_, err = d.device.Volumes(context.TODO(), []string{"block"})

	case 1: // Set Volume
		vol := rand.Intn(101)
		block := _BLOCKS[rand.Intn(len(_BLOCKS))]
		cmd = fmt.Sprintf("Set Volume (%d) on (%s)", vol, block)
		err = d.device.SetVolume(context.TODO(), block, vol)

	case 2: // Get Input
		cmd = "Get Input"
		_, err = d.device.AudioVideoInputs(context.TODO())

	case 3: // Set Input
		input := _INPUTS[rand.Intn(len(_INPUTS))]
		output := _OUTPUTS[rand.Intn(len(_OUTPUTS))]
		cmd = fmt.Sprintf("Set Input (%s) for Output (%s)", input, output)
		err = d.device.SetAudioVideoInput(context.TODO(), output, input)

	case 4: // Get Mutes
		cmd = "Get Mutes"
		_, err = d.device.Mutes(context.TODO(), []string{"block"})

	case 5: // Set Mutes
		mute := _MUTES[rand.Intn(len(_MUTES))]
		cmd = fmt.Sprintf("Set Mute (%s)", mute)
		b := boo()
		cmd = fmt.Sprintf("Set Mute (%s) to (%t)", mute, b)
		err = d.device.SetMute(context.TODO(), mute, b)

	default: // Skip
		cmd = "Skip"
	}

	// If an error occurred
	if err != nil {
		return cmd, fmt.Errorf("Error while trying to issue %s command: %w", cmd, err)
	}

	return cmd, nil
}
