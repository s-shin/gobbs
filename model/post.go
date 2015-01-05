package model

import (
	v "github.com/s-shin/gobbs/validation"
	"time"
)

type Post struct {
	Id        int64
	ThreadId  int64
	Name      string
	Content   string
	CreatedAt time.Time
}

func (p *Post) Validate() v.Error {
	var err v.Error
	if !v.MaxLength(p.Name, 20) {
		err = err.Append(&v.ErrorEntry{"name", v.TooLong})
	}
	if !v.Length(p.Content, 1, 1000) {
		err = err.Append(&v.ErrorEntry{"content", v.InvalidLength})
	}
	return err
}
