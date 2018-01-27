package tvapi_test

import (
	"errors"
	"net/http"

	"github.com/golang/mock/gomock"
	tvapi "github.com/keitax/tv-api"
	"github.com/keitax/tv-api/mock_http"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("PanicHandler", func() {
	var (
		c  *gomock.Controller
		h  http.Handler
		ph http.Handler
	)

	BeforeEach(func() {
		c = gomock.NewController(GinkgoT())
	})
	AfterEach(func() {
		c.Finish()
	})

	JustBeforeEach(func() {
		ph = tvapi.PanicHandler(h)
	})

	Describe("ServeHTTP()", func() {
		var rw *mock_http.MockResponseWriter

		BeforeEach(func() {
			rw = mock_http.NewMockResponseWriter(c)
		})
		AfterEach(func() {
			ph.ServeHTTP(rw, nil)
		})

		Context("with a handler which occurs a panic", func() {
			BeforeEach(func() {
				h = tvapi.Handler(func(w http.ResponseWriter, r *http.Request) {
					panic(errors.New("error message"))
				})
			})

			It("writes about the panic", func() {
				rw.EXPECT().WriteHeader(gomock.Any())
				rw.EXPECT().Write([]byte("Internal Server Error"))
			})
		})
	})
})
