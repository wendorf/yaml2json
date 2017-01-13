package acceptance_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"strings"
)

var _ = Describe("yaml2json", func() {
	var command *exec.Cmd

	sampleYaml := `---
cool: yaml
is: great`

	expectedJSON := `{
  "cool": "yaml",
  "is": "great"
}`

	Describe("file path as an argument", func() {
		BeforeEach(func() {
			yamlFile, _ := ioutil.TempFile("", "")
			yamlFile.Write([]byte(sampleYaml))
			yamlFile.Close()
			command = exec.Command(pathToMain, yamlFile.Name())
		})

		It("prints the file as JSON", func() {
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(session, "5s").Should(gexec.Exit(0))

			Expect(session.Out).To(gbytes.Say(expectedJSON))
		})

		Context("when the file does not exist", func() {
			BeforeEach(func() {
				command = exec.Command(pathToMain, fmt.Sprintf("/tmp/yaml2json/fake/%d", rand.Int()))
			})

			It("errors", func() {
				session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())

				Eventually(session, "5s").Should(gexec.Exit(1))

				Expect(session.Out).To(gbytes.Say("no such file or directory"))
			})
		})
	})

	Describe("YAML via STDIN", func() {
		BeforeEach(func() {
			command = exec.Command(pathToMain)
			command.Stdin = strings.NewReader(sampleYaml)
		})

		It("prints the input as JSON", func() {
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(session, "5s").Should(gexec.Exit(0))

			Expect(session.Out).To(gbytes.Say(expectedJSON))
		})
	})

	Context("when the input is not convertible", func() {
		BeforeEach(func() {
			badYaml := "I AM NOT YAML"

			yamlFile, _ := ioutil.TempFile("", "")
			yamlFile.Write([]byte(badYaml))
			yamlFile.Close()
			command = exec.Command(pathToMain, yamlFile.Name())
		})

		It("errors", func() {
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(session, "5s").Should(gexec.Exit(1))

			Expect(session.Out).To(gbytes.Say("Error: Input cannot be converted"))
		})
	})
})
