package internal_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pursuit/gateway/internal"

	"github.com/stretchr/testify/assert"
)

func TestServeHTTP(t *testing.T) {
	config := make(map[string]string)
	config["/users"] = "userservice"
	s := internal.NewServer(config)

	t.Run("not found", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/healthz", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		s.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("exists", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/users", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		s.ServeHTTP(rr, req)
		assert.NotEqual(t, http.StatusNotFound, rr.Code)
	})
}
