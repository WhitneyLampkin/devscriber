package main

import (
	"flag"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	var defaultInput *userInput

	BeforeEach(func() {
		defaultInput = &userInput{
			template: "readme-template",
			name:     "README",
			imageUrl: "./assets/default_image.png",
		}
	})

	AfterEach(func() {
		// Reset osArgs so they can be parsed again
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	})

	Describe("Get user input", func() {
		Context("when using default arguments", func() {
			It("should return a valid default userInput object", func() {
				os.Args = []string{"cmd", "--template=readme-template", "--name=README", "--imageUrl=./assets/default_image.png"}
				result, err := getUserInput()

				Expect(err).To(BeNil())
				Expect(result.template).To(Equal(defaultInput.template))
				Expect(result.name).To(Equal(defaultInput.name))
				Expect(result.imageUrl).To(Equal(defaultInput.imageUrl))
			})
		})

		Context("when no user arguments are provided", func() {
			It("should also return a valid default userInput object", func() {

			})
		})

		Context("when template is invalid", func() {
			It("should return invalid template error", func() {

			})
		})

		Context("when name is invalid", func() {
			It("should return invalid name error", func() {

			})
		})

		Context("when imageUrl is invalid", func() {
			It("should return invalid imageUrl error", func() {

			})
		})
	})
})
