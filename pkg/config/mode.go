package config

import (
	"encoding/json"
	"errors"
)

var ErrInvalidMode = errors.New("invalid mode")

type Mode string

const (
	HTTP Mode = "http"
	FCGI Mode = "fcgi"
)

// UnmarshalJSON implements json.Unmarshaler.
func (mode *Mode) UnmarshalJSON(data []byte) error {
    var v string
    if err := json.Unmarshal(data, &v); err != nil {
        return err
    }

    switch v {
    case string(HTTP):
        *mode = HTTP
    case string(FCGI):
        *mode = FCGI
    default:
        return ErrInvalidMode
    }

    return nil
}

var _ json.Unmarshaler = (*Mode)(nil)
