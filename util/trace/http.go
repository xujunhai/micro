package trace

import (
	"bytes"
	ot "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"io/ioutil"
	"net/http"
	"strconv"
)

// MaxContentLength is the maximum content length for which we'll read and capture
// the contents of the request body. Anything larger will still be traced but the
// body will not be captured as trace metadata.
const MaxContentLength = 1 << 16

// TracedTransport is a traced HTTP transport that captures spans.
type TracedTransport struct {
	*http.Transport
}

var traceComponent = ot.Tag{Key: string(ext.Component), Value: "http"}

// RoundTrip satisfies the RoundTripper interface, wraps the sub Transport and
// captures a span of the http request.
func (t *TracedTransport) RoundTrip(r *http.Request) (resp *http.Response, err error) {
	span, ctx := ot.StartSpanFromContext(r.Context(), "http.request", traceComponent)
	r = r.WithContext(ctx)
	defer func() {
		if err != nil {
			span.SetTag("http.error", err.Error())
			span.SetTag(string(ext.Error), true)
		}
		span.Finish()
	}()

	span.SetTag(string(ext.HTTPUrl), r.URL.Path)
	span.SetTag("http.method", r.Method)
	span.SetTag("http.url", r.URL.Path)
	span.SetTag("http.params", r.URL.Query().Encode())

	contentLength, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	if r.Body != nil && contentLength < MaxContentLength {
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}

		span.SetTag(string(ext.StringTagName("http.body")), string(buf))
		r.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	}

	err = ot.GlobalTracer().Inject(span.Context(), ot.HTTPHeaders, ot.HTTPHeadersCarrier(r.Header))
	if err != nil {
		return nil, err
	}

	// execute standard roundtrip
	resp, err = t.Transport.RoundTrip(r)
	if err != nil {
		return nil, err
	}

	span.SetTag("http.status_code", resp.StatusCode)
	return resp, err
}

// NewTracedHTTPClient returns a new http.Client with a custom transport
func NewTracedHTTPClient(transport *http.Transport) *http.Client {
	return &http.Client{
		Transport: &TracedTransport{transport},
	}
}
