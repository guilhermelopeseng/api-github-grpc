package server

import (
	"encoding/json"
	"io"
)

func FromJson(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
