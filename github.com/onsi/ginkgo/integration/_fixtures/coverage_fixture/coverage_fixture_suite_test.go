package coverage_fixture_test

import (
	. "github.com/gocircuit/escher/github.com/onsi/ginkgo"
	. "github.com/gocircuit/escher/github.com/onsi/gomega"

	"testing"
)

func TestCoverageFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CoverageFixture Suite")
}
