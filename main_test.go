package main

import (
	"errors"
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
				os.Args = []string{"cmd"}
				result, err := getUserInput()

				Expect(err).To(BeNil())
				Expect(result.template).To(Equal(defaultInput.template))
				Expect(result.name).To(Equal(defaultInput.name))
				Expect(result.imageUrl).To(Equal(defaultInput.imageUrl))
			})
		})

		// TODO: Revisit this functionality when I update the validate template code
		/* Context("when template is invalid", func() {
			It("should return invalid template error", func() {
				expectedErr := errors.New("invalid template value. use readme-template or readme")
				os.Args = []string{"cmd", "--template=f4k3"}
				result, err := getUserInput()

				Expect(err).To(Equal(expectedErr))
				Expect(result).To(Equal(userInput{}))
			})
		}) */

		Context("when name is invalid", func() {
			It("should return invalid name error", func() {
				expectedErr := errors.New("there was an error setting the document's name")
				os.Args = []string{"cmd", "--name="}
				result, err := getUserInput()

				Expect(err).To(Equal(expectedErr))
				Expect(result).To(Equal(userInput{}))
			})
		})

		Context("when imageUrl is invalid", func() {
			It("should return invalid imageUrl error", func() {
				expectedErr := errors.New("there was an error setting the image's Url")
				os.Args = []string{"cmd", "--imageUrl="}
				result, err := getUserInput()

				Expect(err).To(Equal(expectedErr))
				Expect(result).To(Equal(userInput{}))
			})
		})
	})

	Describe("Generate file", func() {
		Context("when the file template is readme-template", func() {
			It("should create a new file with no errors", func() {
				testInput := userInput{
					template: "readme-template",
				}
				result, err := generateFile(testInput.template)

				Expect(err).To(BeNil())
				Expect(result).To(Not(BeNil()))
				// TODO: add check for new file?
			})
		})
	})
})
