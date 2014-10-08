package asyncassertion_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega"

	"testing"
)

func TestAsyncAssertion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AsyncAssertion Suite")
}
