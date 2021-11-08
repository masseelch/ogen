// Code generated by ogen, DO NOT EDIT.

package techempower

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
	"github.com/go-faster/errors"
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

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// Caching implements Caching operation.
	Caching(ctx context.Context, params CachingParams) (WorldObjects, error)
	// DB implements DB operation.
	DB(ctx context.Context) (WorldObject, error)
	// JSON implements json operation.
	JSON(ctx context.Context) (HelloWorld, error)
	// Queries implements Queries operation.
	Queries(ctx context.Context, params QueriesParams) (WorldObjects, error)
	// Updates implements Updates operation.
	Updates(ctx context.Context, params UpdatesParams) (WorldObjects, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	mux *chi.Mux
	cfg config
}

func NewServer(h Handler, opts ...Option) *Server {
	srv := &Server{
		h:   h,
		mux: chi.NewMux(),
		cfg: newConfig(opts...),
	}
	srv.setupRoutes()
	return srv
}

func (s *Server) setupRoutes() {
	s.mux.MethodFunc("GET", "/cached-worlds", s.HandleCachingRequest)
	s.mux.MethodFunc("GET", "/db", s.HandleDBRequest)
	s.mux.MethodFunc("GET", "/json", s.HandleJSONRequest)
	s.mux.MethodFunc("GET", "/queries", s.HandleQueriesRequest)
	s.mux.MethodFunc("GET", "/updates", s.HandleUpdatesRequest)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
