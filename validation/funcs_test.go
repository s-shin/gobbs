package validation

import (
	v "github.com/s-shin/gobbs/validation"
	"testing"
)

func TestMinLength(t *testing.T) {
	seeds := []struct {
		Str    string
		Min    int
		Result bool
	}{
		{"foo", 3, true},
		{"foo", 4, false},
	}
	for _, seed := range seeds {
		if v.MinLength(seed.Str, seed.Min) != seed.Result {
			t.Error(seed)
		}
	}
}

func TestMaxLength(t *testing.T) {
	seeds := []struct {
		Str    string
		Max    int
		Result bool
	}{
		{"foo", 3, true},
		{"foo", 2, false},
	}
	for _, seed := range seeds {
		if v.MaxLength(seed.Str, seed.Max) != seed.Result {
			t.Error(seed)
		}
	}
}

func TestLength(t *testing.T) {
	seeds := []struct {
		Str    string
		Min    int
		Max    int
		Result bool
	}{
		{"foo", 3, 3, true},
		{"foo", 2, 3, true},
		{"foo", 3, 4, true},
		{"foo", 1, 2, false},
		{"foo", 4, 5, false},
	}
	for _, seed := range seeds {
		if v.Length(seed.Str, seed.Min, seed.Max) != seed.Result {
			t.Error(seed)
		}
	}
}
