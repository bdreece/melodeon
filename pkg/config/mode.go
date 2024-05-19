package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Mode int

const (
	HTTP Mode = iota
	FCGI
)

func (m *Mode) UnmarshalYAML(value *yaml.Node) error {
	var s string
	if err := value.Decode(&s); err != nil {
		return err
	}

	switch s {
	case "http":
		*m = HTTP
	case "fcgi":
		*m = FCGI
	default:
		return fmt.Errorf("invalid mode: %q", s)
	}

	return nil
}
