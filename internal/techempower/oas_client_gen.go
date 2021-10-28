// Code generated by ogen, DO NOT EDIT.

package techempower

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	serverURL *url.URL
	cfg       config
	requests  metric.Int64Counter
	errors    metric.Int64Counter
	duration  metric.Int64Histogram
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...Option) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c := &Client{
		cfg:       newConfig(opts...),
		serverURL: u,
	}
	if c.requests, err = c.cfg.Meter.NewInt64Counter(otelogen.ClientRequestCount); err != nil {
		return nil, err
	}
	if c.errors, err = c.cfg.Meter.NewInt64Counter(otelogen.ClientErrorsCount); err != nil {
		return nil, err
	}
	if c.duration, err = c.cfg.Meter.NewInt64Histogram(otelogen.ClientDuration); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Client) Caching(ctx context.Context, params CachingParams) (res WorldObjects, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `Caching`,
		trace.WithAttributes(otelogen.OperationID(`Caching`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/cached-worlds"

	q := u.Query()
	{
		// Encode "count" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		v := params.Count
		param := e.EncodeInt64(v)
		q.Set("count", param)
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeCachingResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) DB(ctx context.Context) (res WorldObject, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `DB`,
		trace.WithAttributes(otelogen.OperationID(`DB`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/db"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeDBResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) JSON(ctx context.Context) (res HelloWorld, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `JSON`,
		trace.WithAttributes(otelogen.OperationID(`json`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/json"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeJSONResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) Queries(ctx context.Context, params QueriesParams) (res WorldObjects, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `Queries`,
		trace.WithAttributes(otelogen.OperationID(`Queries`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/queries"

	q := u.Query()
	{
		// Encode "queries" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		v := params.Queries
		param := e.EncodeInt64(v)
		q.Set("queries", param)
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeQueriesResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) Updates(ctx context.Context, params UpdatesParams) (res WorldObjects, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `Updates`,
		trace.WithAttributes(otelogen.OperationID(`Updates`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/updates"

	q := u.Query()
	{
		// Encode "queries" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		v := params.Queries
		param := e.EncodeInt64(v)
		q.Set("queries", param)
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeUpdatesResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}
