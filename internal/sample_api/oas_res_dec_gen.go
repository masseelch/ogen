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
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
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
)

func decodePetGetResponse(resp *http.Response) (res PetGetResponse, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIterator()
			defer json.PutIterator(i)
			i.ResetBytes(buf.Bytes())

			var response PetGetResponseOKApplicationJSON
			if err := func() error {
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIterator()
			defer json.PutIterator(i)
			i.ResetBytes(buf.Bytes())

			var response PetGetResponseDefaultApplicationJSONStatusCode
			if err := func() error {
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePetCreateResponse(resp *http.Response) (res PetGetResponseOKApplicationJSON, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIterator()
			defer json.PutIterator(i)
			i.ResetBytes(buf.Bytes())

			var response PetGetResponseOKApplicationJSON
			if err := func() error {
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodePetGetByNameResponse(resp *http.Response) (res PetGetResponseOKApplicationJSON, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIterator()
			defer json.PutIterator(i)
			i.ResetBytes(buf.Bytes())

			var response PetGetResponseOKApplicationJSON
			if err := func() error {
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeFoobarGetResponse(resp *http.Response) (res FoobarGetResponse, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIterator()
			defer json.PutIterator(i)
			i.ResetBytes(buf.Bytes())

			var response PetGetResponseOKApplicationJSON
			if err := func() error {
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 404:
		return &FoobarGetResponseNotFound{}, nil
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeFoobarPutResponse(resp *http.Response) (res FoobarPutResponseDefaultStatusCode, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	default:
		return FoobarPutResponseDefaultStatusCode{StatusCode: resp.StatusCode}, nil
	}
}

func decodeFoobarPostResponse(resp *http.Response) (res FoobarPostResponse, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIterator()
			defer json.PutIterator(i)
			i.ResetBytes(buf.Bytes())

			var response PetGetResponseOKApplicationJSON
			if err := func() error {
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 404:
		return &FoobarPostResponseNotFound{}, nil
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIterator()
			defer json.PutIterator(i)
			i.ResetBytes(buf.Bytes())

			var response FoobarPostResponseDefaultApplicationJSONStatusCode
			if err := func() error {
				return nil
			}(); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePetNameByIDResponse(resp *http.Response) (res string, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIterator()
			defer json.PutIterator(i)
			i.ResetBytes(buf.Bytes())

			var response string
			if err := func() error {
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodePetFriendsNamesByIDResponse(resp *http.Response) (res []string, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}

	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			i := json.GetIterator()
			defer json.PutIterator(i)
			i.ResetBytes(buf.Bytes())

			var response []string
			if err := func() error {
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}
