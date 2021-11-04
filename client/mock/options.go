package mock

import (
	"gitlab.ziroom.com/rent-web/micro/client"
)

// Response sets the response methods for a service
func Response(service string, response []MockResponse) client.Option {
	return func(o *client.Options) {
		r, ok := fromContext(o.Context)
		if !ok {
			r = make(map[string][]MockResponse)
		}
		r[service] = response
		o.Context = newContext(o.Context, r)
	}
}
