package uri

type constval struct {
	v string
}

func (v constval) DecodeValue() (string, error) {
	return v.v, nil
}

func (d constval) DecodeArray(f func(Decoder) error) error {
	panic("its a value, not an array")
}

func (d constval) DecodeFields(f func(string, Decoder) error) error {
	panic("its a value, not an object")
}
