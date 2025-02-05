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

func decodeFoobarPostRequest(r *http.Request, span trace.Span) (req OptPet, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		if r.ContentLength == 0 {
			return req, nil
		}
		var request OptPet
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
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
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeOneofBugRequest(r *http.Request, span trace.Span) (req OneOfBugs, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request OneOfBugs
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePetCreateRequest(r *http.Request, span trace.Span) (req OptPet, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		if r.ContentLength == 0 {
			return req, nil
		}
		var request OptPet
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
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
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePetUpdateNameAliasPostRequest(r *http.Request, span trace.Span) (req OptPetName, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		if r.ContentLength == 0 {
			return req, nil
		}
		var request OptPetName
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
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
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePetUpdateNamePostRequest(r *http.Request, span trace.Span) (req OptString, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		if r.ContentLength == 0 {
			return req, nil
		}
		var request OptString
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
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
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePetUploadAvatarByIDRequest(r *http.Request, span trace.Span) (req PetUploadAvatarByIDReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/octet-stream":
		return PetUploadAvatarByIDReq{Data: r.Body}, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}
