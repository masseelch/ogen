// Code generated by ogen, DO NOT EDIT.

package api

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

func (c *Client) FoobarGet(ctx context.Context, params FoobarGetParams) (res FoobarGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `FoobarGet`,
		trace.WithAttributes(otelogen.OperationID(`foobarGet`)),
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
	u.Path += "/foobar"

	q := u.Query()
	{
		// Encode "inlinedParam" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		v := params.InlinedParam
		param := e.EncodeInt64(v)
		q.Set("inlinedParam", param)
	}
	{
		// Encode "skip" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		v := params.Skip
		param := e.EncodeInt32(v)
		q.Set("skip", param)
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeFoobarGetResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) FoobarPost(ctx context.Context, req Pet) (res FoobarPostRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `FoobarPost`,
		trace.WithAttributes(otelogen.OperationID(`foobarPost`)),
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
	buf, contentType, err := encodeFoobarPostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/foobar"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeFoobarPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) FoobarPut(ctx context.Context) (res FoobarPutDefStatusCode, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `FoobarPut`,
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
	u.Path += "/foobar"

	r := ht.NewRequest(ctx, "PUT", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeFoobarPutResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetCreate(ctx context.Context, req PetCreateReq) (res Pet, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PetCreate`,
		trace.WithAttributes(otelogen.OperationID(`petCreate`)),
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
	buf, contentType, err := encodePetCreateRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/pet"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodePetCreateResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetFriendsNamesByID(ctx context.Context, params PetFriendsNamesByIDParams) (res []string, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PetFriendsNamesByID`,
		trace.WithAttributes(otelogen.OperationID(`petFriendsNamesByID`)),
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
	u.Path += "/pet/friendNames/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		u.Path += e.EncodeInt(params.ID)
	}

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodePetFriendsNamesByIDResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetGet(ctx context.Context, params PetGetParams) (res PetGetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PetGet`,
		trace.WithAttributes(otelogen.OperationID(`petGet`)),
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
	u.Path += "/pet"

	q := u.Query()
	{
		// Encode "petID" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		v := params.PetID
		param := e.EncodeInt64(v)
		q.Set("petID", param)
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	{
		value := conv.UUIDArrayToString(params.XTags)
		for _, v := range value {
			r.Header.Add("x-tags", v)
		}
	}
	{
		value := conv.StringArrayToString(params.XScope)
		for _, v := range value {
			r.Header.Add("x-scope", v)
		}
	}

	{
		value := conv.StringToString(params.Token)
		r.AddCookie(&http.Cookie{
			Name:   "token",
			Value:  value,
			MaxAge: 1337,
		})
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodePetGetResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetGetByName(ctx context.Context, params PetGetByNameParams) (res Pet, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PetGetByName`,
		trace.WithAttributes(otelogen.OperationID(`petGetByName`)),
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
	u.Path += "/pet/"
	{
		// Encode "name" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "name",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		u.Path += e.EncodeString(params.Name)
	}

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodePetGetByNameResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetNameByID(ctx context.Context, params PetNameByIDParams) (res string, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PetNameByID`,
		trace.WithAttributes(otelogen.OperationID(`petNameByID`)),
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
	u.Path += "/pet/name/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		u.Path += e.EncodeInt(params.ID)
	}

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodePetNameByIDResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetUpdateNameAliasPost(ctx context.Context, req PetName) (res PetUpdateNameAliasPostDefStatusCode, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PetUpdateNameAliasPost`,
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
	buf, contentType, err := encodePetUpdateNameAliasPostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/pet/updateNameAlias"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodePetUpdateNameAliasPostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetUpdateNamePost(ctx context.Context, req string) (res PetUpdateNamePostDefStatusCode, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PetUpdateNamePost`,
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
	buf, contentType, err := encodePetUpdateNamePostRequest(req, span)
	if err != nil {
		return res, err
	}
	defer json.PutBuffer(buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/pet/updateName"

	r := ht.NewRequest(ctx, "POST", u, buf)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodePetUpdateNamePostResponse(resp, span)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}
