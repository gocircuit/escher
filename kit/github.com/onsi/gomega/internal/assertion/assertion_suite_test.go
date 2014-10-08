package assertion_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega"

	"testing"
)

func TestAssertion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Assertion Suite")
}
