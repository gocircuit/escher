package more_ginkgo_tests_test

import (
	. "github.com/gocircuit/escher/github.com/onsi/ginkgo"
	. "github.com/gocircuit/escher/github.com/onsi/gomega"

	"testing"
)

func TestMore_ginkgo_tests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "More_ginkgo_tests Suite")
}
