package validation

import (
	"bytes"
	"fmt"
)

type Validatable interface {
	Validate() error
}

type ErrorType string

const (
	InvalidLength ErrorType = "invalid length"
	TooLong       ErrorType = "too long"
	TooShort      ErrorType = "too short"
)

func (t ErrorType) String() string {
	return string(t)
}

type ErrorEntry struct {
	Name string
	Type ErrorType
}

func (ee *ErrorEntry) String() string {
	return fmt.Sprintf("%s: %s", ee.Name, ee.Type)
}

type Error []*ErrorEntry

func (e Error) Error() string {
	return e.String()
}

func (e Error) String() string {
	var buf bytes.Buffer
	for _, entry := range e {
		buf.WriteString("{Name: ")
		buf.WriteString(entry.Name)
		buf.WriteString(", ")
		buf.WriteString("Type: ")
		buf.WriteString(string(entry.Type))
		buf.WriteString("} ")
	}
	return buf.String()
}

// Append an entriy to `e.Entries`.
func (e Error) Append(ee ...*ErrorEntry) Error {
	return Error(append(e, ee...))
}

// Create the combined error.
func (e Error) Combine(err Error) Error {
	return Error(append(e, err...))
}
