// View Renderer by html/template
//
//	var r Renderable = NewRenderer().LoadFileFunc(func(filename string) ([]byte, error) {
//		return ioutil.ReadFile(path.Join("view", filename))
//	})
//	r.Render(w, "layout.html:index.html", nil)
//
package util

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

// same as ioutil.ReadFile
type LoadFileFunc func(filename string) ([]byte, error)

type Renderer struct {
	templates map[string]*template.Template
	FuncMap   template.FuncMap
	baseName  string
	loadFile  LoadFileFunc
}

func NewRenderer() *Renderer {
	return &Renderer{
		templates: make(map[string]*template.Template),
		FuncMap: template.FuncMap{
			// {{(.NotStr|string)}}
			"string": func(x interface{}) string {
				return fmt.Sprintf("%d", x)
			},
			// {{add 1 2}}
			"add": func(x, y int) int {
				return x + y
			},
			// {{.Str|default("hoge")}}
			"default": func(defaultVal, x string) string {
				if len(x) == 0 {
					return defaultVal
				}
				return x
			},
		},
		baseName: "BASE",
		loadFile: LoadFileFunc(ioutil.ReadFile),
	}
}

func (r *Renderer) BaseName(name string) *Renderer {
	r.baseName = name
	return r
}

func (r *Renderer) LoadFileFunc(fn func(string) ([]byte, error)) *Renderer {
	r.loadFile = LoadFileFunc(fn)
	return r
}

func (r *Renderer) LoadTemplate(tmplFilesStr string) *template.Template {
	key := tmplFilesStr
	// Load the template from the cache if it exists.
	if r.templates[key] != nil {
		return r.templates[key]
	}
	tmplFiles := strings.Split(tmplFilesStr, ":")
	// See `parseFiles` function (http://golang.org/src/text/template/helper.go).
	var t *template.Template
	for _, tmplFile := range tmplFiles {
		name := filepath.Base(tmplFile)
		if t == nil {
			t = template.New(name).Funcs(r.FuncMap)
		}
		// make associated template
		var tmpl *template.Template
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		// load and parse
		var data string
		if b, err := r.loadFile(tmplFile); err == nil {
			data = string(b)
		} else {
			log.Fatal(err)
		}
		if _, err := tmpl.Parse(data); err != nil {
			log.Fatal(err)
		}
	}
	// Cache it.
	r.templates[key] = t
	return t
}

func (r *Renderer) RenderTemplate(w io.Writer, tmpl *template.Template, data interface{}) {
	err := tmpl.ExecuteTemplate(w, r.baseName, data)
	if err != nil {
		log.Fatal(err)
	}
}

// Implementation of controller.Renderable interface.
// `tmpFilesStr` is the concatenation of template files with ":".
// These template files are passes to `template.ParseFiles(...string)`.
func (r *Renderer) Render(w io.Writer, tmplFilesStr string, data interface{}) {
	tmpl := r.LoadTemplate(tmplFilesStr)
	r.RenderTemplate(w, tmpl, data)
}
