package flags_test

import (
	. "github.com/gocircuit/escher/github.com/onsi/ginkgo"
	. "github.com/gocircuit/escher/github.com/onsi/gomega"

	"testing"
)

func TestFlags(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flags Suite")
}
