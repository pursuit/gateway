package internal

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/rs/cors"
)

type Server struct {
	CORS *cors.Cors
	Urls map[string]string
}

func NewServer(urls map[string]string) *Server {
	s := &Server{}

	s.CORS = cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowCredentials: true,
		MaxAge:           86400,
	})
	s.Urls = urls

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	target, found := s.Urls[r.URL.Path]
	if !found {
		http.NotFound(w, r)
		return
	}

	uri, _ := url.Parse(target)
	httputil.NewSingleHostReverseProxy(uri).ServeHTTP(w, r)
}
