package nested_test

import (
	. "github.com/gocircuit/escher/kit/github.com/onsi/ginkgo"
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega"

	"testing"
)

func TestNested(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nested Suite")
}
