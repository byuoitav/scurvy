package atlona4x1

import "github.com/byuoitav/scurvy"

type Option func(*Atlona4x1)

func WithLogger(l scurvy.Logger) Option {
	return func(d *Atlona4x1) {
		d.logger = l
	}
}
