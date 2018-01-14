package tvapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewApplication makes a new tv-api application instance.
func NewApplication() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	})
	return r
}
