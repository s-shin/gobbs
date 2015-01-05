package controller

import (
	"github.com/gorilla/mux"
	m "github.com/s-shin/gobbs/model"
	s "github.com/s-shin/gobbs/service"
	v "github.com/s-shin/gobbs/validation"
	"log"
	"net/http"
	"strconv"
)

type Post struct {
	*C
}

func getThreadId(vars map[string]string) int64 {
	threadId, err := strconv.ParseInt(vars["thread_id"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return threadId
}

func (c *Post) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	threadId := getThreadId(vars)
	thread := s.FetchThreadById(threadId)
	posts := s.FetchAllThreadPosts(threadId)
	data := struct {
		Thread *m.Thread
		Posts  []*m.Post
	}{
		Thread: thread,
		Posts:  posts,
	}
	c.Render(w, "layout.html:post/index.html", data)
}

func (c *Post) CreateForm(w http.ResponseWriter, r *http.Request) {
	data := new(struct {
		Error v.Error
	})
	c.Render(w, "layout.html:post/create.html", data)
}

func (c *Post) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	threadId := getThreadId(vars)
	name := r.FormValue("post_name")
	content := r.FormValue("post_content")
	_, err := s.CreatePost(threadId, name, content)

	if err != nil {
		if e, ok := err.(v.Error); ok {
			data := struct {
				Error v.Error
			}{e}
			c.Render(w, "layout.html:post/create.html", data)
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
