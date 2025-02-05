package parser

import (
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/internal/oas"
)

type parser struct {
	// schema specification, immutable.
	spec *ogen.Spec
	// parsed operations.
	operations []*oas.Operation
	// refs contains lazy-initialized referenced components.
	refs struct {
		schemas       map[string]*oas.Schema
		requestBodies map[string]*oas.RequestBody
		responses     map[string]*oas.Response
		parameters    map[string]*oas.Parameter
		examples      map[string]*oas.Example
	}
}

func Parse(spec *ogen.Spec) ([]*oas.Operation, error) {
	spec.Init()
	p := &parser{
		spec: spec,
		refs: struct {
			schemas       map[string]*oas.Schema
			requestBodies map[string]*oas.RequestBody
			responses     map[string]*oas.Response
			parameters    map[string]*oas.Parameter
			examples      map[string]*oas.Example
		}{
			schemas:       map[string]*oas.Schema{},
			requestBodies: map[string]*oas.RequestBody{},
			responses:     map[string]*oas.Response{},
			parameters:    map[string]*oas.Parameter{},
			examples:      map[string]*oas.Example{},
		},
	}

	err := p.parse()
	return p.operations, err
}

func (p *parser) parse() error {
	operationIDs := make(map[string]struct{})
	for path, item := range p.spec.Paths {
		if item.Ref != "" {
			return errors.Errorf("%s: referenced pathItem not supported", path)
		}

		itemParams, err := p.parseParams(item.Parameters)
		if err != nil {
			return errors.Wrapf(err, "%s: parameters", path)
		}

		if err := forEachOps(item, func(method string, op ogen.Operation) error {
			if id := op.OperationID; id != "" {
				if _, ok := operationIDs[id]; ok {
					return errors.Errorf("duplicate operationId: %q", id)
				}
				operationIDs[id] = struct{}{}
			}

			parsedOp, err := p.parseOp(path, method, op, itemParams)
			if err != nil {
				return err
			}

			p.operations = append(p.operations, parsedOp)
			return nil
		}); err != nil {
			return errors.Wrapf(err, "paths: %s", path)
		}
	}

	return nil
}

func (p *parser) parseOp(path, httpMethod string, spec ogen.Operation, itemParams []*oas.Parameter) (_ *oas.Operation, err error) {
	op := &oas.Operation{
		OperationID: spec.OperationID,
		Description: spec.Description,
		HTTPMethod:  strings.ToUpper(httpMethod),
	}

	opParams, err := p.parseParams(spec.Parameters)
	if err != nil {
		return nil, errors.Wrap(err, "parameters")
	}

	// Merge operation parameters with pathItem parameters.
	op.Parameters = mergeParams(opParams, itemParams)

	op.Path, err = parsePath(path, op.Parameters)
	if err != nil {
		return nil, errors.Wrap(err, "parse path")
	}

	if spec.RequestBody != nil {
		op.RequestBody, err = p.parseRequestBody(spec.RequestBody)
		if err != nil {
			return nil, errors.Wrap(err, "requestBody")
		}
	}

	op.Responses, err = p.parseResponses(spec.Responses)
	if err != nil {
		return nil, errors.Wrap(err, "responses")
	}

	return op, nil
}

func mergeParams(opParams, itemParams []*oas.Parameter) []*oas.Parameter {
	lookupOp := func(name string, in oas.ParameterLocation) bool {
		for _, param := range opParams {
			if param.Name == name && param.In == in {
				return true
			}
		}
		return false
	}

	for _, param := range itemParams {
		// Param defined in operation take precedense over param defined in pathItem.
		if lookupOp(param.Name, param.In) {
			continue
		}

		opParams = append(opParams, param)
	}

	return opParams
}
