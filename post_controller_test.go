package tvapi_test

import (
	"github.com/golang/mock/gomock"
	tvapi "github.com/keitax/tv-api"
	"github.com/keitax/tv-api/mock_tvapi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PostController", func() {
	var (
		pc tvapi.PostController
		ps *mock_tvapi.MockPostService
	)

	BeforeEach(func() {
		c := gomock.NewController(GinkgoT())
		ps = mock_tvapi.NewMockPostService(c)
		pc = tvapi.NewPostController(ps)
	})

	Describe("GetList()", func() {
		var (
			got     []*tvapi.Post
			err     error
			start   int
			results int
		)

		JustBeforeEach(func() {
			got, err = pc.GetList(start, results)
		})

		Context("with a service which returns some posts", func() {
			BeforeEach(func() {
				start = 1
				results = 2
				ps.EXPECT().GetList(1, 2).AnyTimes().Return([]*tvapi.Post{
					{Key: "post-0"},
					{Key: "post-1"},
				}, nil)
			})

			It("returns those posts", func() {
				Expect(got).To(HaveLen(2))
				Expect(got[0].Key).To(Equal("post-0"))
				Expect(got[1].Key).To(Equal("post-1"))
			})

			It("occurs no error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("GetSingle()", func() {
		var (
			got *tvapi.Post
			err error
			key string
		)

		JustBeforeEach(func() {
			got, err = pc.GetSingle(key)
		})

		Context("with a post", func() {
			BeforeEach(func() {
				key = "20180101-post-key"
				ps.EXPECT().GetSingle("20180101-post-key").AnyTimes().Return(&tvapi.Post{Key: "20180101-post-key"}, nil)
			})

			It("returns the post", func() {
				Expect(got.Key).To(Equal("20180101-post-key"))
			})

			It("occurs no error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Synchronize()", func() {
		var err error

		JustBeforeEach(func() {
			err = pc.Synchronize()
		})

		Context("with a service expected to be synchronized", func() {
			BeforeEach(func() {
				ps.EXPECT().Synchronize().Times(1).Return(nil)
			})

			It("occurs no error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
