package gen

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ogen-go/ogen"
)

type methodResponse struct {
	Responses map[int]*Response
	Default   *Response
}

func (g *Generator) generateResponses(methodName string, methodResponses ogen.Responses) (*methodResponse, error) {
	var (
		responses   = make(map[int]*Response)
		defaultResp *Response
	)

	// Iterate over method responses...
	for status, schema := range methodResponses {
		// Default response.
		if status == "default" {
			resp, err := g.createDefaultResponse(methodName, schema)
			if err != nil {
				return nil, fmt.Errorf("default: %w", err)
			}

			defaultResp = resp
			continue
		}

		statusCode, err := strconv.Atoi(status)
		if err != nil {
			return nil, fmt.Errorf("invalid status code: '%s'", status)
		}

		if err := func() error {
			// Referenced response.
			if ref := schema.Ref; ref != "" {
				// Validate reference & get response component name.
				name, err := componentName(ref)
				if err != nil {
					return err
				}

				// Lookup for response component.
				r, found := g.responses[name]
				if !found {
					return fmt.Errorf("response by reference '%s' not found", ref)
				}

				responses[statusCode] = r
				return nil
			}

			responseName := pascal(methodName, http.StatusText(statusCode))
			resp, err := g.generateResponse(responseName, schema)
			if err != nil {
				return err
			}

			responses[statusCode] = resp
			return nil
		}(); err != nil {
			return nil, fmt.Errorf("%s: %w", status, err)
		}
	}

	return &methodResponse{
		Responses: responses,
		Default:   defaultResp,
	}, nil
}

// createDefaultResponse creates new default response.
func (g *Generator) createDefaultResponse(methodName string, r ogen.Response) (*Response, error) {
	if ref := r.Ref; ref != "" {
		// Validate reference & get response component name.
		name, err := componentName(ref)
		if err != nil {
			return nil, err
		}

		// Lookup for alias response.
		if alias, ok := g.responses[name+"Default"]; ok {
			return alias, nil
		}

		// Lookup for reference response.
		response, found := g.responses[name]
		if !found {
			return nil, fmt.Errorf("response by reference '%s', not found", ref)
		}

		alias := g.createResponse()
		for contentType, schema := range response.Contents {
			alias.Contents[contentType] = g.wrapStatusCode(schema)
		}
		if schema := response.NoContent; schema != nil {
			response.NoContent = g.wrapStatusCode(schema)
		}

		g.responses[name+"Default"] = alias
		return alias, nil
	}

	// Inlined response.
	// Use method name + Default as prefix for response schemas.
	response, err := g.generateResponse(methodName+"Default", r)
	if err != nil {
		return nil, err
	}

	// We need to inject StatusCode field to response structs somehow...
	// Iterate over all responses and create new response schema wrapper:
	//
	// type <WrapperName> struct {
	//     StatusCode int            `json:"-"`
	//     Response   <ResponseType> `json:"-"`
	// }
	for contentType, schema := range response.Contents {
		defaultSchema := g.wrapStatusCode(schema)
		response.Contents[contentType] = defaultSchema
	}
	if schema := response.NoContent; schema != nil {
		response.NoContent = g.wrapStatusCode(schema)
	}

	return response, nil
}

// generateResponse creates new response based on schema definition.
func (g *Generator) generateResponse(rname string, resp ogen.Response) (*Response, error) {
	response := g.createResponse()

	// Response without content.
	// Create empty struct.
	if len(resp.Content) == 0 {
		s := g.createSchemaAlias(rname, "struct{}")
		g.schemas[s.Name] = s
		response.NoContent = s
		return response, nil
	}

	for contentType, media := range resp.Content {
		// Create unique response name.
		name := rname + "Response"
		if len(resp.Content) > 1 {
			name = pascal(rname, contentType, "Response")
		}

		// Referenced response schema.
		if ref := media.Schema.Ref; ref != "" {
			refSchemaName, err := componentName(ref)
			if err != nil {
				return nil, fmt.Errorf("content: %s: %w", contentType, err)
			}

			schema, found := g.schemas[refSchemaName]
			if !found {
				return nil, fmt.Errorf("content: %s: schema referenced by '%s' not found", contentType, ref)
			}

			// Response have only one content.
			// Use schema directly without creating new one.
			if len(resp.Content) == 1 {
				response.Contents[contentType] = schema
				continue
			}

			// Response have multiple contents.
			// Alias them with new response type.
			s := g.createSchemaAlias(name, schema.Name)
			g.schemas[s.Name] = s
			response.Contents[contentType] = s
			continue
		}

		// Inlined response schema.
		s, err := g.generateSchema(name, media.Schema)
		if err != nil {
			return nil, fmt.Errorf("content: %s: schema: %w", contentType, err)
		}

		g.schemas[s.Name] = s
		response.Contents[contentType] = s
	}

	return response, nil
}

// wrapStatusCode wraps provided schema with newtype containing StatusCode field.
//
// Example 1:
//   Schema:
//   type FoobarGetDefaultResponse {
//       Message string `json:"message"`
//       Code    int64  `json:"code"`
//   }
//
//   Wrapper:
//   type FoobarGetDefaultResponseStatusCode {
//       StatusCode int                      `json:"-"`
//       Response   FoobarGetDefaultResponse `json:"-"`
//   }
//
// Example 2:
//   Schema:
//   type FoobarGetDefaultResponse string
//
//   Wrapper:
//   type FoobarGetDefaultResponseStatusCode {
//       StatusCode int    `json:"-"`
//       Response   string `json:"-"`
//   }
//
// TODO: Remove unused schema (Example 2).
func (g *Generator) wrapStatusCode(schema *Schema) *Schema {
	// Use 'StatusCode' postfix for wrapper struct name
	// to avoid name collision with original response schema.
	newSchema := g.createSchemaStruct(schema.Name + "StatusCode")
	newSchema.Fields = []SchemaField{
		{
			Name: "StatusCode",
			Tag:  "-",
			Type: "int",
		},
		{
			Name: "Response",
			Tag:  "-",
			Type: schema.Type(),
		},
	}
	g.schemas[newSchema.Name] = newSchema
	return newSchema
}
