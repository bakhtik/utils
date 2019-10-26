// Package json contains json wrappers
package json

import (
	js "encoding/json"
	"net/url"
)

func Marshal(v interface{}, values url.Values) ([]byte, error) {
	if _, ok := values["debug"]; ok {
		return js.MarshalIndent(v, "", "   ")
	}

	return js.Marshal(v)
}
