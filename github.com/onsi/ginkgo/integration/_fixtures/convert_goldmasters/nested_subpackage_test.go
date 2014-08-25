package subpackage

import (
	. "github.com/gocircuit/escher/github.com/onsi/ginkgo"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("nested sub packages", func() {
		GinkgoT().Fail(true)
	})
})
