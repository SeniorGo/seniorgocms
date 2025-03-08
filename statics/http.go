package statics

import (
	"embed"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"strings"
)

// Serve static files
//
//go:embed www/*
var www embed.FS

func FileReader(statics string) func(filename string) ([]byte, error) {
	return func(filename string) ([]byte, error) {
		if statics == "" {
			return www.ReadFile(path.Join("www", filename))
		}
		return os.ReadFile(path.Join(statics, filename))
	}
}

func ServeStatics(statics string) http.HandlerFunc {
	if statics == "" {
		return AddPrefix("../www", http.FileServer(http.FS(www)))
	}
	if strings.HasPrefix(statics, "http://") || strings.HasPrefix(statics, "https://") {

		director := func(rr *http.Request) {
			u, _ := url.Parse(statics)
			rr.Host = u.Host
			rr.URL.Scheme = u.Scheme
			rr.URL.Host = u.Host
			rr.URL.Path = u.Path + strings.TrimPrefix(rr.URL.Path, u.Path)
		}

		proxy := httputil.ReverseProxy{
			FlushInterval: 0,
			// FlushInterval:  50 * time.Millisecond, // -1 will flush immediately
			Director: director,
			// Transport:      proxyTransport,
			// ModifyResponse: modifyResponse,
		}

		return proxy.ServeHTTP
	}
	return http.FileServer(http.Dir(statics)).ServeHTTP
}

// Copied from http.StripPrefix
func AddPrefix(prefix string, h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := prefix + r.URL.Path
		rp := prefix + r.URL.Path
		r2 := new(http.Request)
		*r2 = *r
		r2.URL = new(url.URL)
		*r2.URL = *r.URL
		r2.URL.Path = p
		r2.URL.RawPath = rp
		h.ServeHTTP(w, r2)
	}
}
