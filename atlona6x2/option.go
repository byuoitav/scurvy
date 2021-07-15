package atlona6x2

import "github.com/byuoitav/scurvy"

type Option func(*Atlona6x2)

func WithLogger(l scurvy.Logger) Option {
	return func(d *Atlona6x2) {
		d.logger = l
	}
}
