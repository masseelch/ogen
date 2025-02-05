package ir

import (
	"fmt"
	"strings"
	"unicode"
)

func capitalize(s string) string {
	var v []rune
	for i, c := range s {
		if i == 0 {
			c = unicode.ToUpper(c)
		}
		v = append(v, c)
	}
	return string(v)
}

func afterDot(v string) string {
	idx := strings.Index(v, ".")
	if idx > 0 {
		return v[idx+1:]
	}
	return v
}

func (t *Type) EncodeFn() string {
	if t.Is(KindArray) && t.Item.EncodeFn() != "" {
		return t.Item.EncodeFn() + "Array"
	}
	switch t.Primitive {
	case Int, Int64, Int32, String, Bool, Float32, Float64:
		return capitalize(t.Primitive.String())
	case UUID, Time, IP, Duration, URL:
		return afterDot(t.Primitive.String())
	default:
		return ""
	}
}

func (t Type) ToString() string {
	if t.EncodeFn() == "" {
		return ""
	}
	return t.EncodeFn() + "ToString"
}

func (t Type) FromString() string {
	if t.EncodeFn() == "" {
		return ""
	}
	return "To" + t.EncodeFn()
}

func (t *Type) IsInteger() bool {
	switch t.Primitive {
	case Int, Int8, Int16, Int32, Int64,
		Uint, Uint8, Uint16, Uint32, Uint64:
		return true
	default:
		return false
	}
}

func (t *Type) IsFloat() bool {
	switch t.Primitive {
	case Float32, Float64:
		return true
	default:
		return false
	}
}

func (t *Type) IsArray() bool     { return t.Is(KindArray) }
func (t *Type) IsMap() bool       { return t.Is(KindMap) }
func (t *Type) IsPrimitive() bool { return t.Is(KindPrimitive) }
func (t *Type) IsStruct() bool    { return t.Is(KindStruct) }
func (t *Type) IsPointer() bool   { return t.Is(KindPointer) }
func (t *Type) IsEnum() bool      { return t.Is(KindEnum) }
func (t *Type) IsGeneric() bool   { return t.Is(KindGeneric) }
func (t *Type) IsAlias() bool     { return t.Is(KindAlias) }
func (t *Type) IsInterface() bool { return t.Is(KindInterface) }
func (t *Type) IsSum() bool       { return t.Is(KindSum) }
func (t *Type) IsStream() bool    { return t.Is(KindStream) }
func (t *Type) IsNumeric() bool   { return t.IsInteger() || t.IsFloat() }

func (t *Type) ReceiverType() string {
	if t.needsPointerReceiverType() {
		return "*" + t.Name
	}
	return t.Name
}

// AdditionalPropsField returns name of additional properties field.
func (t *Type) AdditionalPropsField() string {
	return "AdditionalProps"
}

func (t *Type) needsPointerReceiverType() bool {
	switch t.Kind {
	case KindPointer, KindArray, KindMap:
		return false
	case KindAlias:
		return t.AliasTo.needsPointerReceiverType()
	default:
		return true
	}
}

func (t *Type) MustField(name string) *Field {
	if !t.Is(KindStruct) {
		panic("unreachable")
	}

	for _, f := range t.Fields {
		if f.Name == name {
			return f
		}
	}

	panic(fmt.Sprintf("field with name %q not found", name))
}
