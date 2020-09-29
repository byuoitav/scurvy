package adcp

import "github.com/byuoitav/scurvy"

type Option func(*RandomDispatcher)

func WithInputs(inputs ...string) Option {
	return func(d *RandomDispatcher) {
		d.inputs = inputs
	}
}

func WithLogger(l scurvy.Logger) Option {
	return func(d *RandomDispatcher) {
		d.logger = l
	}
}
