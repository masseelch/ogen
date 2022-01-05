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

	args := map[string]string{}
	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if prefix := "/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
				elem = elem[len(prefix):]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleErrorGetRequest(args, w, r)
				return
			}
			switch elem[0] {
			case 'e': // Prefix: "error"
				if prefix := "error"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				// Leaf: ErrorGet
				s.handleErrorGetRequest(args, w, r)
				return
			case 'f': // Prefix: "foobar"
				if prefix := "foobar"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				// Leaf: FoobarGet
				s.handleFoobarGetRequest(args, w, r)
				return
			case 'n': // Prefix: "name/"
				if prefix := "name/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				// Param: "id"
				// Match until one of "/"
				idx := strings.IndexByte(elem, '/')
				if idx > 0 {
					args["id"] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if prefix := "/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Param: "foo"
						// Match until one of "1"
						idx := strings.IndexByte(elem, '1')
						if idx > 0 {
							args["foo"] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '1': // Prefix: "1234"
								if prefix := "1234"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
									elem = elem[len(prefix):]
								} else {
									break
								}

								// Param: "bar"
								// Match until one of "-"
								idx := strings.IndexByte(elem, '-')
								if idx > 0 {
									args["bar"] = elem[:idx]
									elem = elem[idx:]

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case '-': // Prefix: "-"
										if prefix := "-"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
											elem = elem[len(prefix):]
										} else {
											break
										}

										// Param: "baz"
										// Match until one of "!"
										idx := strings.IndexByte(elem, '!')
										if idx > 0 {
											args["baz"] = elem[:idx]
											elem = elem[idx:]

											if len(elem) == 0 {
												break
											}
											switch elem[0] {
											case '!': // Prefix: "!"
												if prefix := "!"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
													elem = elem[len(prefix):]
												} else {
													break
												}

												// Param: "kek"
												// Leaf parameter
												args["kek"] = elem

												// Leaf: DataGetFormat
												s.handleDataGetFormatRequest(args, w, r)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			case 'p': // Prefix: "pet"
				if prefix := "pet"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handlePetGetRequest(args, w, r)
					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if prefix := "/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handlePetGetAvatarByIDRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'a': // Prefix: "avatar"
						if prefix := "avatar"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: PetGetAvatarByID
						s.handlePetGetAvatarByIDRequest(args, w, r)
						return
					case 'f': // Prefix: "friendNames/"
						if prefix := "friendNames/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Param: "id"
						// Leaf parameter
						args["id"] = elem

						// Leaf: PetFriendsNamesByID
						s.handlePetFriendsNamesByIDRequest(args, w, r)
						return
					case 'n': // Prefix: "name/"
						if prefix := "name/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Param: "id"
						// Leaf parameter
						args["id"] = elem

						// Leaf: PetNameByID
						s.handlePetNameByIDRequest(args, w, r)
						return
					}
					// Param: "name"
					// Leaf parameter
					args["name"] = elem

					// Leaf: PetGetByName
					s.handlePetGetByNameRequest(args, w, r)
					return
				}
			case 't': // Prefix: "test/header"
				if prefix := "test/header"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				// Leaf: GetHeader
				s.handleGetHeaderRequest(args, w, r)
				return
			}
		}
	case "POST":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if prefix := "/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
				elem = elem[len(prefix):]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handlePetCreateRequest(args, w, r)
				return
			}
			switch elem[0] {
			case 'f': // Prefix: "foobar"
				if prefix := "foobar"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				// Leaf: FoobarPost
				s.handleFoobarPostRequest(args, w, r)
				return
			case 'p': // Prefix: "pet"
				if prefix := "pet"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handlePetCreateRequest(args, w, r)
					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if prefix := "/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
						elem = elem[len(prefix):]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handlePetUploadAvatarByIDRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'a': // Prefix: "avatar"
						if prefix := "avatar"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						// Leaf: PetUploadAvatarByID
						s.handlePetUploadAvatarByIDRequest(args, w, r)
						return
					case 'u': // Prefix: "updateName"
						if prefix := "updateName"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
							elem = elem[len(prefix):]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handlePetUpdateNamePostRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'A': // Prefix: "Alias"
							if prefix := "Alias"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
								elem = elem[len(prefix):]
							} else {
								break
							}

							// Leaf: PetUpdateNameAliasPost
							s.handlePetUpdateNameAliasPostRequest(args, w, r)
							return
						}
					}
				}
			}
		}
	case "PUT":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/foobar"
			if prefix := "/foobar"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
				elem = elem[len(prefix):]
			} else {
				break
			}

			// Leaf: FoobarPut
			s.handleFoobarPutRequest(args, w, r)
			return
		}
	}
	s.notFound(w, r)
}
