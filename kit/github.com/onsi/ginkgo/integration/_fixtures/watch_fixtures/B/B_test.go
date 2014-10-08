package B_test

import (
	. "github.com/gocircuit/escher/kit/github.com/onsi/ginkgo/integration/_fixtures/watch_fixtures/B"

	. "github.com/gocircuit/escher/kit/github.com/onsi/ginkgo"
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega"
)

var _ = Describe("B", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
