package validation

import (
	v "github.com/s-shin/gobbs/validation"
	"testing"
)

func makeError() v.Error {
	return v.Error{
		&v.ErrorEntry{"foo", v.TooLong},
		&v.ErrorEntry{"bar", v.TooShort},
	}
}

func TestErrorAppend(t *testing.T) {
	var e v.Error
	e = e.Append(&v.ErrorEntry{}, &v.ErrorEntry{})
	expected := 2
	actual := len(e)
	if expected != actual {
		t.Errorf("Append %d entries but the length is %d.", expected, actual)
	}
}

func TestCombineError(t *testing.T) {
	e := makeError().Combine(makeError())
	expected := 4
	actual := len(e)
	if expected != actual {
		t.Errorf("Combine 2 errors both of that contain 2 entries, but the length is %d.", actual)
	}
}
