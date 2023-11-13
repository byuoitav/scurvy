package vscontrol

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/byuoitav/scurvy"
)

type VSControl struct {
	logger scurvy.Logger
	r      *rand.Rand
	device *VSController
}

type NullLogger struct{}

func (l *NullLogger) Debugf(format string, a ...interface{}) {}
func (l *NullLogger) Infof(format string, a ...interface{})  {}
func (l *NullLogger) Warnf(format string, a ...interface{})  {}
func (l *NullLogger) Errorf(format string, a ...interface{}) {}

func New(address, port string, encoders, decoders *[]string) scurvy.RandomDispatcher {
	d := VSControl{
		device: &VSController{
			address:     address,
			port:        port,
			encoderList: encoders,
			decoderList: decoders,
		},
	}
	d.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return &d
}

func (vs *VSControl) RandomDispatch() (cmd string, err error) {

	switch vs.r.Intn(5) {
	case 0:
		cmd = "Set Host"
		_, err = vs.device.SetStreamHost()
	case 1:
		cmd = "Set Video Wall"
		_, err = vs.device.SetVideoWall()
	case 2:
		cmd = "Check Connection"
		_, err = vs.device.GetConnection()
	case 3:
		cmd = "Get Device Info"
		_, err = vs.device.GetDeviceInfo()
	case 4:
		cmd = "Get Signal"
		_, err = vs.device.GetSignal()
	default:
		cmd = "Skip"
	}

	if err != nil {
		return cmd, fmt.Errorf("failure while trying to issue %s command: %w", cmd, err)
	}

	return cmd, nil
}
