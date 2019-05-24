package ghttp_test

import (
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"

	"testing"
)

func TestGHTTP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GHTTP Suite")
}
