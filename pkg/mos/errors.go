package mos

import "fmt"

// EmulationError an error occurred at emulation.
type EmulationError struct {
	// Type is the error's type.
	Type string
	// Desc is the error's description.
	Desc string
}

func (e *EmulationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Type, e.Desc)
}

func NewEmulationError(t string, f string, a ...any) error {
	return &EmulationError{
		Type: t,
		Desc: f,
	}
}

// IsError reports if the given error is a certain error type.
func IsError(t string, err error) bool {
	if e, ok := err.(*EmulationError); !ok {
		return false
	} else {
		return e.Type == t
	}
}
