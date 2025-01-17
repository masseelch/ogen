package parser

import (
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/internal/oas"
)

// schemaParser is used to parse OpenAPI schemas.
type schemaParser struct {
	// Schema components from OpenAPI document (readonly).
	components map[string]*ogen.Schema

	// Parsed schema components from *parser (readonly).
	// Used as cache from previous parsing ops.
	globalRefs map[string]*oas.Schema

	// Parsed schema components current schema refers to.
	localRefs map[string]*oas.Schema
}

func (p *schemaParser) Parse(schema *ogen.Schema) (*oas.Schema, error) {
	return p.parse(schema, func(s *oas.Schema) *oas.Schema {
		return p.extendInfo(schema, s)
	})
}

func (p *schemaParser) parse(schema *ogen.Schema, hook func(*oas.Schema) *oas.Schema) (*oas.Schema, error) {
	if schema == nil {
		return nil, nil
	}

	if ref := schema.Ref; ref != "" {
		s, err := p.resolve(ref)
		if err != nil {
			return nil, errors.Wrapf(err, "reference %q", ref)
		}
		return s, nil
	}

	switch {
	case len(schema.Enum) > 0:
		if err := validateTypeFormat(schema.Type, schema.Format); err != nil {
			return nil, errors.Wrap(err, "validate format")
		}

		values, err := parseEnumValues(oas.SchemaType(schema.Type), schema.Enum)
		if err != nil {
			return nil, errors.Wrap(err, "parse enum")
		}

		return hook(&oas.Schema{
			Type:   oas.SchemaType(schema.Type),
			Format: oas.Format(schema.Format),
			Enum:   values,
		}), nil
	case len(schema.OneOf) > 0:
		schemas, err := p.parseMany(schema.OneOf)
		if err != nil {
			return nil, errors.Wrapf(err, "oneOf")
		}

		return hook(&oas.Schema{OneOf: schemas}), nil
	case len(schema.AnyOf) > 0:
		schemas, err := p.parseMany(schema.AnyOf)
		if err != nil {
			return nil, errors.Wrapf(err, "anyOf")
		}

		return hook(&oas.Schema{AnyOf: schemas}), nil
	case len(schema.AllOf) > 0:
		schemas, err := p.parseMany(schema.AllOf)
		if err != nil {
			return nil, errors.Wrapf(err, "allOf")
		}

		return hook(&oas.Schema{AllOf: schemas}), nil
	}

	// Try to derive schema type from properties.
	if schema.Type == "" {
		switch {
		case schema.Items != nil:
			schema.Type = "array"
		}
	}

	switch schema.Type {
	case "object":
		if schema.Items != nil {
			return nil, errors.New("object cannot contain 'items' field")
		}
		required := func(name string) bool {
			for _, p := range schema.Required {
				if p == name {
					return true
				}
			}
			return false
		}

		var item *oas.Schema
		if ap := schema.AdditionalProperties; ap != nil {
			sch := ogen.Schema(*ap)
			prop, err := p.Parse(&sch)
			if err != nil {
				return nil, errors.Wrapf(err, "additionalProperties")
			}
			item = prop
		}

		s := hook(&oas.Schema{
			Type:          oas.Object,
			Item:          item,
			MaxProperties: schema.MaxProperties,
			MinProperties: schema.MinProperties,
		})
		for _, propSpec := range schema.Properties {
			prop, err := p.Parse(propSpec.Schema)
			if err != nil {
				return nil, errors.Wrapf(err, "%s", propSpec.Name)
			}

			s.Properties = append(s.Properties, oas.Property{
				Name:        propSpec.Name,
				Description: propSpec.Schema.Description,
				Schema:      prop,
				Required:    required(propSpec.Name),
			})
		}
		return s, nil

	case "array":
		array := hook(&oas.Schema{
			Type:        oas.Array,
			MinItems:    schema.MinItems,
			MaxItems:    schema.MaxItems,
			UniqueItems: schema.UniqueItems,
		})

		if len(schema.Properties) > 0 {
			return nil, errors.New("array cannot contain properties")
		}

		item, err := p.Parse(schema.Items)
		if err != nil {
			return nil, errors.Wrap(err, "item")
		}

		array.Item = item
		return array, nil

	case "number", "integer":
		if err := validateTypeFormat(schema.Type, schema.Format); err != nil {
			return nil, errors.Wrap(err, "validate format")
		}

		return hook(&oas.Schema{
			Type:             oas.SchemaType(schema.Type),
			Format:           oas.Format(schema.Format),
			Minimum:          schema.Minimum,
			Maximum:          schema.Maximum,
			ExclusiveMinimum: schema.ExclusiveMinimum,
			ExclusiveMaximum: schema.ExclusiveMaximum,
			MultipleOf:       schema.MultipleOf,
		}), nil

	case "boolean":
		if err := validateTypeFormat(schema.Type, schema.Format); err != nil {
			return nil, errors.Wrap(err, "validate format")
		}

		return hook(&oas.Schema{
			Type:   oas.Boolean,
			Format: oas.Format(schema.Format),
		}), nil

	case "string":
		if err := validateTypeFormat(schema.Type, schema.Format); err != nil {
			return nil, errors.Wrap(err, "validate format")
		}

		return hook(&oas.Schema{
			Type:      oas.String,
			Format:    oas.Format(schema.Format),
			MaxLength: schema.MaxLength,
			MinLength: schema.MinLength,
			Pattern:   schema.Pattern,
		}), nil

	case "":
		return nil, nil

	default:
		return nil, errors.Errorf("unexpected schema type: %q", schema.Type)
	}
}

func (p *schemaParser) parseMany(schemas []*ogen.Schema) ([]*oas.Schema, error) {
	result := make([]*oas.Schema, 0, len(schemas))
	for i, schema := range schemas {
		s, err := p.Parse(schema)
		if err != nil {
			return nil, errors.Wrapf(err, "[%d]", i)
		}

		result = append(result, s)
	}

	return result, nil
}

func (p *schemaParser) resolve(ref string) (*oas.Schema, error) {
	const prefix = "#/components/schemas/"
	if !strings.HasPrefix(ref, prefix) {
		return nil, errors.Errorf("invalid schema reference %q", ref)
	}

	if s, ok := p.globalRefs[ref]; ok {
		return s, nil
	}

	if s, ok := p.localRefs[ref]; ok {
		return s, nil
	}

	name := strings.TrimPrefix(ref, prefix)
	component, found := p.components[name]
	if !found {
		return nil, errors.Errorf("component by reference %q not found", ref)
	}

	return p.parse(component, func(s *oas.Schema) *oas.Schema {
		s.Ref = ref
		p.localRefs[ref] = s
		return p.extendInfo(component, s)
	})
}

func (p *schemaParser) extendInfo(schema *ogen.Schema, s *oas.Schema) *oas.Schema {
	s.Description = schema.Description
	s.AddExample(schema.Example)

	// Workaround: handle nullable enums correctly.
	//
	// Notice that nullable enum requires `null` in value list.
	//
	// See https://github.com/OAI/OpenAPI-Specification/blob/main/proposals/2019-10-31-Clarify-Nullable.md#if-a-schema-specifies-nullable-true-and-enum-1-2-3-does-that-schema-allow-null-values-see-1900.
	if len(s.Enum) < 1 {
		s.Nullable = schema.Nullable
	} else {
		// Check that enum contains `null` value.
		for _, v := range s.Enum {
			if v == nil {
				s.Nullable = true
				break
			}
		}
		// Filter all `null`s.
		if s.Nullable {
			n := 0
			for _, v := range s.Enum {
				if v != nil {
					s.Enum[n] = v
					n++
				}
			}
			s.Enum = s.Enum[:n]
		}
	}
	if d := schema.Discriminator; d != nil {
		s.Discriminator = &oas.Discriminator{
			PropertyName: d.PropertyName,
			Mapping:      d.Mapping,
		}
	}
	return s
}

func validateTypeFormat(typ, format string) error {
	formats := map[string]map[string]struct{}{
		"object":  {},
		"array":   {},
		"boolean": {},
		"integer": {
			"int32": {},
			"int64": {},
		},
		"number": {
			"float":  {},
			"double": {},
			"int32":  {},
			"int64":  {},
		},
		"string": {
			"byte":      {},
			"date-time": {},
			"date":      {},
			"time":      {},
			"duration":  {},
			"uuid":      {},
			"ipv4":      {},
			"ipv6":      {},
			"ip":        {},
			"uri":       {},
			"password":  {},
		},
	}

	if _, ok := formats[typ]; !ok {
		return errors.Errorf("unexpected type: %q", typ)
	}

	if format == "" {
		return nil
	}

	if _, ok := formats[typ][format]; !ok {
		if typ == "string" {
			return nil // Ignore unknown string formats.
		}

		return errors.Errorf("unexpected %s format: %q", typ, format)
	}

	return nil
}
