package atlona5x2

import "github.com/byuoitav/scurvy"

type Option func(*Atlona5x2)

func WithLogger(l scurvy.Logger) Option {
	return func(d *Atlona5x2) {
		d.logger = l
	}
}
