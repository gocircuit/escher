package format_test

import (
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"

	"testing"
)

func TestFormat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Format Suite")
}
