package decoder

import (
	"encoding/json"
	"io"
)

func DecodeJSON(r io.Reader, v interface{}) error {
	defer func() { _, _ = io.Copy(io.Discard, r) }()
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}

func EncodeJSON(w io.Writer, v interface{}) error {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(true)
	return encoder.Encode(v)
}
