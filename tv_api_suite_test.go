package tvapi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTvApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TvApi Suite")
}
