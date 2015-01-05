package model

import (
	v "github.com/s-shin/gobbs/validation"
	"time"
)

type Thread struct {
	Id          int64
	Title       string
	DefaultName string
	PostNum     int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Thread) Validate() v.Error {
	var err v.Error
	if !v.Length(t.Title, 1, 100) {
		err = err.Append(&v.ErrorEntry{"title", v.InvalidLength})
	}
	if !v.MaxLength(t.DefaultName, 20) {
		err = err.Append(&v.ErrorEntry{"default name", v.TooLong})
	}
	return err
}
