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

// Server handles operations described by OpenAPI v3 specification.
type Server interface {
	// FoobarGet implements foobarGet operation.
	FoobarGet(ctx context.Context, params FoobarGetParams) (FoobarGetRes, error)
	// FoobarPost implements foobarPost operation.
	FoobarPost(ctx context.Context, req Pet) (FoobarPostRes, error)
	// FoobarPut implements  operation.
	FoobarPut(ctx context.Context) (FoobarPutDefStatusCode, error)
	// PetCreate implements petCreate operation.
	PetCreate(ctx context.Context, req PetCreateReq) (Pet, error)
	// PetFriendsNamesByID implements petFriendsNamesByID operation.
	PetFriendsNamesByID(ctx context.Context, params PetFriendsNamesByIDParams) ([]string, error)
	// PetGet implements petGet operation.
	PetGet(ctx context.Context, params PetGetParams) (PetGetRes, error)
	// PetGetByName implements petGetByName operation.
	PetGetByName(ctx context.Context, params PetGetByNameParams) (Pet, error)
	// PetNameByID implements petNameByID operation.
	PetNameByID(ctx context.Context, params PetNameByIDParams) (string, error)
	// PetUpdateNameAliasPost implements  operation.
	PetUpdateNameAliasPost(ctx context.Context, req PetName) (PetUpdateNameAliasPostDefStatusCode, error)
	// PetUpdateNamePost implements  operation.
	PetUpdateNamePost(ctx context.Context, req string) (PetUpdateNamePostDefStatusCode, error)
}
