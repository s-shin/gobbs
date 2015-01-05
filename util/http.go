package util

import (
	"net"
	"net/http"
	"os"
	"path"
)

type Request struct {
	*http.Request
}

// This function supports a reverse proxy.
func (r *Request) GetIP() string {
	ip := r.Header.Get("X-FORWARDED-FOR")
	if len(ip) == 0 {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

// http.FileServer without indexing.
func StaticServer(dir string) http.Handler {
	fileServer := http.FileServer(http.Dir(dir))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filepath := path.Join(dir, r.URL.Path)
		finfo, err := os.Stat(filepath)
		if err != nil || finfo.Mode().IsDir() {
			http.NotFoundHandler().ServeHTTP(w, r)
		} else {
			fileServer.ServeHTTP(w, r)
		}
	})
}
