package tvapi

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// NewApplication makes a new tv-api application instance.
func NewApplication() http.Handler {
	h := mux.NewRouter()
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	})
	return PanicHandler(h)
}

// Handler is wrapper of handler functions.
type Handler func(http.ResponseWriter, *http.Request)

// ServeHTTP serves requests.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

// PanicHandler wraps handlers to handle panics.
func PanicHandler(h http.Handler) http.Handler {
	return Handler(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				log.Println("Internal Server Error")
			}
		}()
		h.ServeHTTP(w, r)
	})
}
