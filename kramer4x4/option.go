package kramer4x4

import "github.com/byuoitav/scurvy"

type Option func(*Kramer4x4)

func WithLogger(l scurvy.Logger) Option {
	return func(d *Kramer4x4) {
		d.logger = l
	}
}
