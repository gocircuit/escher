package focused_fixture_test

import (
	. "github.com/gocircuit/escher/github.com/onsi/ginkgo"
	. "github.com/gocircuit/escher/github.com/onsi/gomega"

	"testing"
)

func TestFocused_fixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Focused_fixture Suite")
}
