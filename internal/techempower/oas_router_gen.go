// Code generated by ogen, DO NOT EDIT.

package techempower

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

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleDBRequest([0]string{}, w, r)

				return
			}
			switch elem[0] {
			case 'c': // Prefix: "cached-worlds"
				if l := len("cached-worlds"); len(elem) >= l && elem[0:l] == "cached-worlds" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: Caching
					s.handleCachingRequest([0]string{}, w, r)

					return
				}
			case 'd': // Prefix: "db"
				if l := len("db"); len(elem) >= l && elem[0:l] == "db" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: DB
					s.handleDBRequest([0]string{}, w, r)

					return
				}
			case 'j': // Prefix: "json"
				if l := len("json"); len(elem) >= l && elem[0:l] == "json" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: JSON
					s.handleJSONRequest([0]string{}, w, r)

					return
				}
			case 'q': // Prefix: "queries"
				if l := len("queries"); len(elem) >= l && elem[0:l] == "queries" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: Queries
					s.handleQueriesRequest([0]string{}, w, r)

					return
				}
			case 'u': // Prefix: "updates"
				if l := len("updates"); len(elem) >= l && elem[0:l] == "updates" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: Updates
					s.handleUpdatesRequest([0]string{}, w, r)

					return
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name  string
	count int
	args  [0]string
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.name
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
func (s *Server) FindRoute(method, path string) (r Route, _ bool) {
	var (
		args = [0]string{}
		elem = path
	)
	r.args = args

	// Static code generated router with unwrapped path search.
	switch method {
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				r.name = "DB"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'c': // Prefix: "cached-worlds"
				if l := len("cached-worlds"); len(elem) >= l && elem[0:l] == "cached-worlds" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: Caching
					r.name = "Caching"
					r.args = args
					r.count = 0
					return r, true
				}
			case 'd': // Prefix: "db"
				if l := len("db"); len(elem) >= l && elem[0:l] == "db" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: DB
					r.name = "DB"
					r.args = args
					r.count = 0
					return r, true
				}
			case 'j': // Prefix: "json"
				if l := len("json"); len(elem) >= l && elem[0:l] == "json" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: JSON
					r.name = "JSON"
					r.args = args
					r.count = 0
					return r, true
				}
			case 'q': // Prefix: "queries"
				if l := len("queries"); len(elem) >= l && elem[0:l] == "queries" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: Queries
					r.name = "Queries"
					r.args = args
					r.count = 0
					return r, true
				}
			case 'u': // Prefix: "updates"
				if l := len("updates"); len(elem) >= l && elem[0:l] == "updates" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: Updates
					r.name = "Updates"
					r.args = args
					r.count = 0
					return r, true
				}
			}
		}
	}
	return r, false
}
