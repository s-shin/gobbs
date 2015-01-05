// Controller Package
package controller

import (
	"github.com/gorilla/mux"
	"io"
)

type Renderable interface {
	Render(w io.Writer, target string, data interface{})
}

type C struct {
	Router   *mux.Router
	Renderer Renderable
}

func (c *C) Render(w io.Writer, target string, data interface{}) {
	c.Renderer.Render(w, target, data)
}
