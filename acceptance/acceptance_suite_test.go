package acceptance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"
	"testing"
)

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance Suite")
}

var (
	pathToMain string
)

var _ = BeforeSuite(func() {
	var err error
	pathToMain, err = gexec.Build("github.com/wendorf/yaml2json")
	Expect(err).NotTo(HaveOccurred())
})
