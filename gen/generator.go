package gen

import (
	"regexp"
	"sort"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/internal/ir"
	"github.com/ogen-go/ogen/internal/oas"
	"github.com/ogen-go/ogen/internal/oas/parser"
)

type responses struct {
	types     map[string]*ir.Type
	responses map[string]*ir.StatusResponse
}

type Generator struct {
	opt        Options
	operations []*ir.Operation
	types      map[string]*ir.Type
	interfaces map[string]*ir.Type
	refs       responses
	wrapped    responses
	errType    *ir.StatusResponse
	router     Router
}

type Options struct {
	VerboseRoute         bool
	GenerateExampleTests bool
	SkipTestRegex        *regexp.Regexp
	SpecificMethodPath   string
	IgnoreNotImplemented []string
}

func NewGenerator(spec *ogen.Spec, opts Options) (*Generator, error) {
	operations, err := parser.Parse(spec)
	if err != nil {
		return nil, errors.Wrap(err, "parse")
	}

	g := &Generator{
		opt:        opts,
		types:      map[string]*ir.Type{},
		interfaces: map[string]*ir.Type{},
		refs: responses{
			types:     map[string]*ir.Type{},
			responses: map[string]*ir.StatusResponse{},
		},
		wrapped: responses{
			types:     map[string]*ir.Type{},
			responses: map[string]*ir.StatusResponse{},
		},
	}

	if err := g.makeIR(operations); err != nil {
		return nil, errors.Wrap(err, "make ir")
	}
	for _, w := range g.wrapped.types {
		g.saveType(w)
	}
	g.reduce()
	g.wrapGenerics()
	if err := g.route(); err != nil {
		return nil, errors.Wrap(err, "route")
	}
	return g, nil
}

func (g *Generator) makeIR(ops []*oas.Operation) error {
	if err := g.reduceDefault(ops); err != nil {
		return errors.Wrap(err, "reduce default")
	}

	for _, spec := range ops {
		if p := g.opt.SpecificMethodPath; p != "" {
			if spec.Path.String() != p {
				continue
			}
		}

		op, err := g.generateOperation(spec)
		if err != nil {
			if err := g.fail(err); err != nil {
				return errors.Wrapf(err, "%q: %s",
					spec.Path.String(),
					strings.ToLower(spec.HTTPMethod),
				)
			}

			continue
		}
		g.operations = append(g.operations, op)
	}

	sort.SliceStable(g.operations, func(i, j int) bool {
		a, b := g.operations[i], g.operations[j]
		return strings.Compare(a.Name, b.Name) < 0
	})

	return nil
}
