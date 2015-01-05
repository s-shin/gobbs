package main

import (
	"flag"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/s-shin/gobbs/bindata"
	c "github.com/s-shin/gobbs/controller"
	"github.com/s-shin/gobbs/db"
	"github.com/s-shin/gobbs/util"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var DIR = util.CWD()

func newRenderer(router *mux.Router) *util.Renderer {
	r := util.NewRenderer().LoadFileFunc(func(filename string) ([]byte, error) {
		return bindata.Asset(path.Join("tmpl", filename))
	})
	r.FuncMap["title"] = func(name string) string {
		return strings.Title(name) + " | "
	}
	r.FuncMap["url"] = func(name string, params ...string) string {
		url, err := router.Get(name).URL(params...)
		if err != nil {
			log.Panic(err)
		}
		return url.String()
	}
	return r
}

func newRouter() *mux.Router {
	staticDir := path.Join(DIR, "static")

	r := mux.NewRouter()
	mc := &c.C{r, newRenderer(r)}

	// static files
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", util.StaticServer(staticDir)))

	{
		cc := &c.Thread{mc}
		r.HandleFunc("/", cc.Index).Methods("GET").Name("root")
		s := r.PathPrefix("/threads").Subrouter()
		s.HandleFunc("/", cc.Index).Methods("GET").Name("list_threads")
		s.HandleFunc("/create", cc.CreateForm).Methods("GET").Name("create_thread")
		s.HandleFunc("/create", cc.Create).Methods("POST").Name("create_thread")
	}
	{
		cc := &c.Post{mc}
		s := r.PathPrefix("/threads/{thread_id:[0-9]+}/posts").Subrouter()
		s.HandleFunc("/", cc.Show).Methods("GET").Name("show_thread_posts")
		s.HandleFunc("/create", cc.CreateForm).Methods("GET").Name("create_thread_post")
		s.HandleFunc("/create", cc.Create).Methods("POST").Name("create_thread_post")
	}

	return r
}

func beforeHandler(w http.ResponseWriter, r *http.Request) {
}

func afterHandler(w http.ResponseWriter, r *http.Request) {
	db.DisconnectAll()
}

func main() {
	address := flag.String("address", ":8080", "Server address")
	flag.Parse()
	http.Handle("/", newRouter())
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		beforeHandler(w, r)
		h := handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)
		h.ServeHTTP(w, r)
		afterHandler(w, r)
	})
	err := gracehttp.Serve(&http.Server{Addr: *address, Handler: handler})
	if err != nil {
		log.Fatal(err)
	}
}
