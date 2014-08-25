package tmp

import (
	. "github.com/gocircuit/escher/github.com/onsi/ginkgo"
	. "github.com/gocircuit/escher/github.com/onsi/gomega"

	"testing"
)

func TestTmp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tmp Suite")
}
