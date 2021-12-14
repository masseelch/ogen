// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
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

func decodeCreateSnapshotRequest(r *http.Request, span trace.Span) (req SnapshotCreateParams, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request SnapshotCreateParams
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
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeCreateSyncActionRequest(r *http.Request, span trace.Span) (req InstanceActionInfo, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request InstanceActionInfo
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
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeLoadSnapshotRequest(r *http.Request, span trace.Span) (req SnapshotLoadParams, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request SnapshotLoadParams
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

func decodeMmdsConfigPutRequest(r *http.Request, span trace.Span) (req MmdsConfig, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request MmdsConfig
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

func decodeMmdsPatchRequest(r *http.Request, span trace.Span) (req OptMmdsPatchReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		if r.ContentLength == 0 {
			return req, nil
		}
		var request OptMmdsPatchReq
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
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeMmdsPutRequest(r *http.Request, span trace.Span) (req OptMmdsPutReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		if r.ContentLength == 0 {
			return req, nil
		}
		var request OptMmdsPutReq
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
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePatchBalloonRequest(r *http.Request, span trace.Span) (req BalloonUpdate, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request BalloonUpdate
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

func decodePatchBalloonStatsIntervalRequest(r *http.Request, span trace.Span) (req BalloonStatsUpdate, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request BalloonStatsUpdate
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

func decodePatchGuestDriveByIDRequest(r *http.Request, span trace.Span) (req PartialDrive, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request PartialDrive
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
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePatchGuestNetworkInterfaceByIDRequest(r *http.Request, span trace.Span) (req PartialNetworkInterface, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request PartialNetworkInterface
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
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePatchMachineConfigurationRequest(r *http.Request, span trace.Span) (req OptMachineConfiguration, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		if r.ContentLength == 0 {
			return req, nil
		}
		var request OptMachineConfiguration
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

func decodePatchVmRequest(r *http.Request, span trace.Span) (req VM, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request VM
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
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePutBalloonRequest(r *http.Request, span trace.Span) (req Balloon, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Balloon
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

func decodePutGuestBootSourceRequest(r *http.Request, span trace.Span) (req BootSource, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request BootSource
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

func decodePutGuestDriveByIDRequest(r *http.Request, span trace.Span) (req Drive, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Drive
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
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePutGuestNetworkInterfaceByIDRequest(r *http.Request, span trace.Span) (req NetworkInterface, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request NetworkInterface
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
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePutGuestVsockRequest(r *http.Request, span trace.Span) (req Vsock, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Vsock
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
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePutLoggerRequest(r *http.Request, span trace.Span) (req Logger, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Logger
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
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePutMachineConfigurationRequest(r *http.Request, span trace.Span) (req OptMachineConfiguration, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		if r.ContentLength == 0 {
			return req, nil
		}
		var request OptMachineConfiguration
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

func decodePutMetricsRequest(r *http.Request, span trace.Span) (req Metrics, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Metrics
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
