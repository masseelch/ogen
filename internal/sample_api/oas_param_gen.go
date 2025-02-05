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

type DataGetFormatParams struct {
	ID  int
	Foo string
	Bar string
	Baz string
	Kek string
}

type FoobarGetParams struct {
	// InlinedParam.
	InlinedParam int64
	// Number of items to skip.
	Skip int32
}

type GetHeaderParams struct {
	XAuthToken string
}

type PetFriendsNamesByIDParams struct {
	// Pet ID.
	ID int
}

type PetGetParams struct {
	// ID of pet.
	PetID int64
	// Tags of pets.
	XTags []uuid.UUID
	// Pet scopes.
	XScope []string
	// Token.
	Token string
}

type PetGetAvatarByIDParams struct {
	// ID of pet.
	PetID int64
}

type PetGetByNameParams struct {
	// Name of pet.
	Name string
}

type PetNameByIDParams struct {
	// Pet ID.
	ID int
}

type PetUploadAvatarByIDParams struct {
	// ID of pet.
	PetID int64
}

type TestObjectQueryParameterParams struct {
	FormObject OptTestObjectQueryParameterFormObject
	DeepObject OptTestObjectQueryParameterDeepObject
}
