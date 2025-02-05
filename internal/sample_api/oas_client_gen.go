// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
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
	_ = bits.LeadingZeros64
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

// Client implements OAS client.
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

// DataGetFormat invokes dataGetFormat operation.
//
// Retrieve data.
//
// GET /name/{id}/{foo}1234{bar}-{baz}!{kek}
func (c *Client) DataGetFormat(ctx context.Context, params DataGetFormatParams) (res string, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `DataGetFormat`,
		trace.WithAttributes(otelogen.OperationID(`dataGetFormat`)),
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
	u.Path += "/name/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}
	u.Path += "/"
	{
		// Encode "foo" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "foo",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Foo))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}
	u.Path += "1234"
	{
		// Encode "bar" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "bar",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Bar))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}
	u.Path += "-"
	{
		// Encode "baz" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "baz",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Baz))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}
	u.Path += "!"
	{
		// Encode "kek" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "kek",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Kek))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeDataGetFormatResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ErrorGet invokes errorGet operation.
//
// Returns error.
//
// GET /error
func (c *Client) ErrorGet(ctx context.Context) (res ErrorStatusCode, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `ErrorGet`,
		trace.WithAttributes(otelogen.OperationID(`errorGet`)),
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
	u.Path += "/error"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeErrorGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// FoobarGet invokes foobarGet operation.
//
// Dumb endpoint for testing things.
//
// GET /foobar
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
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.InlinedParam))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["inlinedParam"] = e.Result()
	}
	{
		// Encode "skip" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int32ToString(params.Skip))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["skip"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeFoobarGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// FoobarPost invokes foobarPost operation.
//
// Dumb endpoint for testing things.
//
// POST /foobar
func (c *Client) FoobarPost(ctx context.Context, request OptPet) (res FoobarPostRes, err error) {
	if err := func() error {
		if request.Set {
			if err := func() error {
				if err := request.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
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
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeFoobarPostRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/foobar"

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeFoobarPostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// FoobarPut invokes  operation.
//
// PUT /foobar
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
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeFoobarPutResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetHeader invokes getHeader operation.
//
// Test for header param.
//
// GET /test/header
func (c *Client) GetHeader(ctx context.Context, params GetHeaderParams) (res Hash, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `GetHeader`,
		trace.WithAttributes(otelogen.OperationID(`getHeader`)),
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
	u.Path += "/test/header"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	{
		e := uri.NewHeaderEncoder(uri.HeaderEncoderConfig{
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.XAuthToken))
		}(); err != nil {
			return res, errors.Wrap(err, `encode header param x-auth-token`)
		}
		if v, ok := e.Result(); ok {
			r.Header.Set("x-auth-token", v)
		}
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeGetHeaderResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OneofBug invokes oneofBug operation.
//
// POST /oneofBug
func (c *Client) OneofBug(ctx context.Context, request OneOfBugs) (res OneofBugOK, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `OneofBug`,
		trace.WithAttributes(otelogen.OperationID(`oneofBug`)),
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
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeOneofBugRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/oneofBug"

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOneofBugResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PetCreate invokes petCreate operation.
//
// Creates pet.
//
// POST /pet
func (c *Client) PetCreate(ctx context.Context, request OptPet) (res Pet, err error) {
	if err := func() error {
		if request.Set {
			if err := func() error {
				if err := request.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
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
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePetCreateRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/pet"

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePetCreateResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PetFriendsNamesByID invokes petFriendsNamesByID operation.
//
// Returns names of all friends of pet.
//
// GET /pet/friendNames/{id}
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
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePetFriendsNamesByIDResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PetGet invokes petGet operation.
//
// Returns pet from the system that the user has access to.
//
// GET /pet
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
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.PetID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["petID"] = e.Result()
	}
	{
		// Encode "token" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Token))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["token"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	{
		e := uri.NewHeaderEncoder(uri.HeaderEncoderConfig{
			Explode: false,
		})
		if err := func() error {
			return e.EncodeArray(func(e uri.Encoder) error {
				for i, item := range params.XTags {
					if err := func() error {
						return e.EncodeValue(conv.UUIDToString(item))
					}(); err != nil {
						return errors.Wrapf(err, "[%d]", i)
					}
				}
				return nil
			})
		}(); err != nil {
			return res, errors.Wrap(err, `encode header param x-tags`)
		}
		if v, ok := e.Result(); ok {
			r.Header.Set("x-tags", v)
		}
	}
	{
		e := uri.NewHeaderEncoder(uri.HeaderEncoderConfig{
			Explode: false,
		})
		if err := func() error {
			return e.EncodeArray(func(e uri.Encoder) error {
				for i, item := range params.XScope {
					if err := func() error {
						return e.EncodeValue(conv.StringToString(item))
					}(); err != nil {
						return errors.Wrapf(err, "[%d]", i)
					}
				}
				return nil
			})
		}(); err != nil {
			return res, errors.Wrap(err, `encode header param x-scope`)
		}
		if v, ok := e.Result(); ok {
			r.Header.Set("x-scope", v)
		}
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePetGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PetGetAvatarByID invokes petGetAvatarByID operation.
//
// Returns pet avatar by id.
//
// GET /pet/avatar
func (c *Client) PetGetAvatarByID(ctx context.Context, params PetGetAvatarByIDParams) (res PetGetAvatarByIDRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PetGetAvatarByID`,
		trace.WithAttributes(otelogen.OperationID(`petGetAvatarByID`)),
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
	u.Path += "/pet/avatar"

	q := u.Query()
	{
		// Encode "petID" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.PetID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["petID"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePetGetAvatarByIDResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PetGetByName invokes petGetByName operation.
//
// Returns pet by name from the system that the user has access to.
//
// GET /pet/{name}
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
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Name))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePetGetByNameResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PetNameByID invokes petNameByID operation.
//
// Returns pet name by pet id.
//
// GET /pet/name/{id}
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
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePetNameByIDResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PetUpdateNameAliasPost invokes  operation.
//
// POST /pet/updateNameAlias
func (c *Client) PetUpdateNameAliasPost(ctx context.Context, request OptPetName) (res PetUpdateNameAliasPostDefStatusCode, err error) {
	if err := func() error {
		if request.Set {
			if err := func() error {
				if err := request.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
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
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePetUpdateNameAliasPostRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/pet/updateNameAlias"

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePetUpdateNameAliasPostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PetUpdateNamePost invokes  operation.
//
// POST /pet/updateName
func (c *Client) PetUpdateNamePost(ctx context.Context, request OptString) (res PetUpdateNamePostDefStatusCode, err error) {
	if err := func() error {
		if request.Set {
			if err := func() error {
				if err := (validate.String{
					MinLength:    6,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(request.Value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
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
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodePetUpdateNamePostRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/pet/updateName"

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePetUpdateNamePostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PetUploadAvatarByID invokes petUploadAvatarByID operation.
//
// Uploads pet avatar by id.
//
// POST /pet/avatar
func (c *Client) PetUploadAvatarByID(ctx context.Context, request PetUploadAvatarByIDReq, params PetUploadAvatarByIDParams) (res PetUploadAvatarByIDRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `PetUploadAvatarByID`,
		trace.WithAttributes(otelogen.OperationID(`petUploadAvatarByID`)),
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
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/octet-stream"
	buf, err := encodePetUploadAvatarByIDRequestOctetStream(request, span)
	if err != nil {
		return res, err
	}
	reqBody = buf

	u := uri.Clone(c.serverURL)
	u.Path += "/pet/avatar"

	q := u.Query()
	{
		// Encode "petID" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.PetID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["petID"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePetUploadAvatarByIDResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TestObjectQueryParameter invokes testObjectQueryParameter operation.
//
// GET /testObjectQueryParameter
func (c *Client) TestObjectQueryParameter(ctx context.Context, params TestObjectQueryParameterParams) (res TestObjectQueryParameterOK, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `TestObjectQueryParameter`,
		trace.WithAttributes(otelogen.OperationID(`testObjectQueryParameter`)),
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
	u.Path += "/testObjectQueryParameter"

	q := u.Query()
	{
		// Encode "formObject" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.FormObject.Get(); ok {
				return val.encodeURI(e)
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["formObject"] = e.Result()
	}
	{
		// Encode "deepObject" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleDeepObject,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.DeepObject.Get(); ok {
				return val.encodeURI(e)
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["deepObject"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeTestObjectQueryParameterResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
