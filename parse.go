package ogen

import (
	"encoding/json"

	"github.com/go-faster/errors"

	"github.com/go-faster/jx"
	"github.com/goccy/go-yaml"
)

func Parse(data []byte) (*Spec, error) {
	if !jx.Valid(data) {
		d, err := yaml.YAMLToJSON(data)
		if err != nil {
			return nil, err
		}
		data = d
	}
	if len(data) == 0 {
		return nil, errors.New("blank data")
	}

	s := &Spec{}
	if err := json.Unmarshal(data, s); err != nil {
		return nil, err
	}

	return s, nil
}
