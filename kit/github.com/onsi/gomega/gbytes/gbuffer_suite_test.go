package gbytes_test

import (
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"

	"testing"
)

func TestGbytes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gbytes Suite")
}
