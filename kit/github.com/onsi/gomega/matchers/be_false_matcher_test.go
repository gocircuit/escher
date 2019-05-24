package matchers_test

import (
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega"
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega/matchers"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("BeFalse", func() {
	It("should handle true and false correctly", func() {
		Ω(true).ShouldNot(BeFalse())
		Ω(false).Should(BeFalse())
	})

	It("should only support booleans", func() {
		success, err := (&BeFalseMatcher{}).Match("foo")
		Ω(success).Should(BeFalse())
		Ω(err).Should(HaveOccurred())
	})
})
