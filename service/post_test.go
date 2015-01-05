package service

import (
	"github.com/s-shin/gobbs/db"
	s "github.com/s-shin/gobbs/service"
	"github.com/s-shin/gobbs/util"
	"path"
	"testing"
)

func init_() {
	if err := db.InitDB(); err != nil {
		panic(err)
	}
	p := path.Join(util.ProjectDir(), "script", "sql", "fixtures", "post_test.sql")
	if err := db.QueryFile(db.Master(), p); err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	init_()
	postId, err := s.CreatePost(1, "bar", "bar")
	if err != nil {
		panic(err)
	}
	expected := int64(4)
	if postId != expected {
		t.Errorf("Last insert post id should be %d, but %d.", expected, postId)
	}
}

func TestFetchAllThreadPosts(t *testing.T) {
	init_()
	posts := s.FetchAllThreadPosts(2)
	expected := 2
	if len(posts) != expected {
		t.Errorf("%d posts should be selected, but %d.", expected, len(posts))
	}
}
