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
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/errors"
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

func decodeAnswerCallbackQueryPostResponse(resp *http.Response, span trace.Span) (res AnswerCallbackQueryPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response AnswerCallbackQueryPostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeAnswerPreCheckoutQueryPostResponse(resp *http.Response, span trace.Span) (res AnswerPreCheckoutQueryPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response AnswerPreCheckoutQueryPostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeAnswerShippingQueryPostResponse(resp *http.Response, span trace.Span) (res AnswerShippingQueryPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response AnswerShippingQueryPostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeClosePostResponse(resp *http.Response, span trace.Span) (res ClosePostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ClosePostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeDeleteStickerFromSetPostResponse(resp *http.Response, span trace.Span) (res DeleteStickerFromSetPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response DeleteStickerFromSetPostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeDeleteWebhookPostResponse(resp *http.Response, span trace.Span) (res DeleteWebhookPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response DeleteWebhookPostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetFilePostResponse(resp *http.Response, span trace.Span) (res GetFilePostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response GetFilePostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetGameHighScoresPostResponse(resp *http.Response, span trace.Span) (res GetGameHighScoresPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response GetGameHighScoresPostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetMePostResponse(resp *http.Response, span trace.Span) (res GetMePostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response GetMePostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetMyCommandsPostResponse(resp *http.Response, span trace.Span) (res GetMyCommandsPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response GetMyCommandsPostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetStickerSetPostResponse(resp *http.Response, span trace.Span) (res GetStickerSetPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response GetStickerSetPostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetUpdatesPostResponse(resp *http.Response, span trace.Span) (res GetUpdatesPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response GetUpdatesPostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetUserProfilePhotosPostResponse(resp *http.Response, span trace.Span) (res GetUserProfilePhotosPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response GetUserProfilePhotosPostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetWebhookInfoPostResponse(resp *http.Response, span trace.Span) (res GetWebhookInfoPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response GetWebhookInfoPostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeLogOutPostResponse(resp *http.Response, span trace.Span) (res LogOutPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response LogOutPostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeSendGamePostResponse(resp *http.Response, span trace.Span) (res SendGamePostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response SendGamePostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeSendInvoicePostResponse(resp *http.Response, span trace.Span) (res SendInvoicePostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response SendInvoicePostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeSetMyCommandsPostResponse(resp *http.Response, span trace.Span) (res SetMyCommandsPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response SetMyCommandsPostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeSetStickerPositionInSetPostResponse(resp *http.Response, span trace.Span) (res SetStickerPositionInSetPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response SetStickerPositionInSetPostOK
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := json.GetBuffer()
			defer json.PutBuffer(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := json.GetDecoder()
			defer json.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response ErrorStatusCode
			if err := func() error {
				if err := response.ReadJSON(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}
