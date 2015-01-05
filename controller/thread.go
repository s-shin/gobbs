package controller

import (
	m "github.com/s-shin/gobbs/model"
	s "github.com/s-shin/gobbs/service"
	v "github.com/s-shin/gobbs/validation"
	"log"
	"net/http"
	"strconv"
)

type Thread struct {
	*C
}

func (c *Thread) Index(w http.ResponseWriter, r *http.Request) {
	data := new(struct {
		Threads []*m.Thread
	})
	data.Threads = s.FetchRecentThreads(0, 11)
	c.Render(w, "layout.html:thread/index.html", data)
}

func (c *Thread) CreateForm(w http.ResponseWriter, r *http.Request) {
	data := new(struct {
		Error v.Error
	})
	c.Render(w, "layout.html:thread/create.html", data)
}

func (c *Thread) Create(w http.ResponseWriter, r *http.Request) {
	threadTitle := r.FormValue("thread_title")
	threadDefaultName := r.FormValue("thread_default_name")
	postName := r.FormValue("post_name")
	postContent := r.FormValue("post_content")
	threadId, err := s.CreateThread(threadTitle, threadDefaultName, postName, postContent)

	if err != nil {
		if e, ok := err.(v.Error); ok {
			data := struct {
				Error v.Error
			}{e}
			c.Render(w, "layout.html:thread/create.html", data)
			return
		}
		log.Fatal(err)
	}

	if url, err := c.Router.Get("show_thread_posts").URL("thread_id", strconv.FormatInt(threadId, 10)); err != nil {
		log.Fatal(err)
	} else {
		http.Redirect(w, r, url.String(), 302)
	}
}
