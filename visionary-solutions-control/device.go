package vscontrol

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/byuoitav/scurvy"
)

type VSController struct {
	address     string
	port        string
	logger      scurvy.Logger
	encoderList *[]string
	decoderList *[]string
}

func NewController(address, port string, encoders, decoders *[]string, logger scurvy.Logger) *VSController {
	return &VSController{
		address:     address,
		port:        port,
		logger:      logger,
		encoderList: encoders,
		decoderList: decoders,
	}
}

func (vs *VSController) getEncoder() string {
	index := rand.Intn(len(*vs.encoderList))
	return (*vs.encoderList)[index]
}

func (vs *VSController) getDecoder() string {
	index := rand.Intn(len(*vs.decoderList))
	return (*vs.decoderList)[index]
}

func (vs *VSController) SetStreamHost() (string, error) {
	resp, err := vs.sendRequest("", "GET", "input/"+vs.getEncoder()+"/"+vs.getDecoder())
	if err != nil {
		return "", fmt.Errorf("failed to set host: %w", err)
	}

	return resp, nil
}

func (vs *VSController) SetVideoWall() (string, error) {
	return "", nil
}

func (vs *VSController) GetConnection() (string, error) {
	resp, err := vs.sendRequest("", "GET", "input/get/"+vs.getEncoder())
	if err != nil {
		return "", fmt.Errorf("failed to get encoder address: %w", err)
	}

	return resp, nil
}

func (vs *VSController) GetDeviceInfo() (string, error) {
	resp, err := vs.sendRequest("", "GET", vs.getDecoder()+"/hardware")
	if err != nil {
		return "", fmt.Errorf("failed to get device info: %w", err)
	}

	return resp, nil
}

func (vs *VSController) GetSignal() (string, error) {
	resp, err := vs.sendRequest("", "GET", vs.getDecoder()+"/signal")
	if err != nil {
		return "", fmt.Errorf("failed to get signal: %w", err)
	}

	return resp, nil
}

func (vs *VSController) sendRequest(body, method, endpoint string) (string, error) {

	req, err := http.NewRequest(method, fmt.Sprintf("http://%s:%s/%s", vs.address, vs.port, endpoint), bytes.NewReader([]byte(body)))
	if err != nil {
		return "", fmt.Errorf("failed to make request")
	}

	client := http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return "", fmt.Errorf("request failed. non-200 response")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body")
	}

	return string(b), nil
}
