package util

import (
	"github.com/s-shin/gobbs/util"
	"io/ioutil"
	"path"
	"strings"
	"testing"
)

func newRenderer() *util.Renderer {
	return util.NewRenderer().LoadFileFunc(func(filename string) ([]byte, error) {
		return ioutil.ReadFile(path.Join(util.CallerDir(), "fixtures", "renderer", filename))
	})
}

type stringWriter struct {
	str string
}

func (w *stringWriter) Write(p []byte) (int, error) {
	w.str += string(p)
	return len(p), nil
}
func (w *stringWriter) String() string {
	return w.str
}

func TestRender(t *testing.T) {
	r := newRenderer()
	w := new(stringWriter)
	r.Render(w, "layout.html:index.html", nil)
	if len(w.String()) == 0 {
		t.Errorf("No string is rendered.")
	}
}

func TestFuncMap(t *testing.T) {
	r := newRenderer()
	r.FuncMap["concat"] = func(s1, s2 string) string {
		return s1 + s2
	}
	w := new(stringWriter)
	r.Render(w, "layout.html:funcmap.html", nil)
	s := w.String()
	if len(s) == 0 {
		t.Errorf("No string is rendered.")
	}
	strs := []string{
		"<p>3</p>",
		"<p>100200</p>",
		"<p>foobar</p>",
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("Should contain %s.", str)
		}
	}
}
