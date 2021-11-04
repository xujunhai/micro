package random

import (
	"gitlab.ziroom.com/rent-web/micro/selector"
	"math/rand"
)

type random struct{}

func (r *random) Select(routes []string, opts ...selector.SelectOption) (selector.Next, error) {
	// we can't select from an empty pool of routes
	if len(routes) == 0 {
		return nil, selector.ErrNoneAvailable
	}

	// return the next func
	return func() string {
		// if there is only one route provided we'll select it
		if len(routes) == 1 {
			return routes[0]
		}

		// select a random route from the slice [0,n)
		return routes[rand.Intn(len(routes))]
	}, nil
}

func (r *random) Record(addr string, err error) error {
	return nil
}

func (r *random) Reset() error {
	return nil
}

func (r *random) String() string {
	return "random"
}

// NewSelector returns a random selector
func NewSelector(opts ...selector.Option) selector.Selector {
	return new(random)
}
