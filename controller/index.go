package controller

import (
	"net/http"
)

type Root struct {
	*C
}

func (c *Root) Index(w http.ResponseWriter, r *http.Request) {
	c.Render(w, "layout.html:index.html", nil)
}
