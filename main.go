package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// A userInput stores the options provided by the user.
type userInput struct {
	template string
	name     string
	imageUrl string
}

func main() {
	// Help info displayed when the --help flag is provided
	flag.Usage = func() {
		fmt.Printf("Usage: %s [OPTIONS]...\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Get user input
	inputs, err := getUserInput()
	if err != nil {
		exitGracefully(err)
	}

	fmt.Println(inputs)
}

// Gets user input from the terminal and validates it
// Returns a validated userInput struct
func getUserInput() (userInput, error) {
	// Validate the number of required arguments
	if len(os.Args) == 0 {
		return userInput{}, errors.New("the template argument is required")
	}

	// Define flags for the template, name and imageUrl arguments
	template := flag.String("template", "readme-template", "Template type to base the new document on")
	name := flag.String("name", "README.md", "Name of the new document")
	imageUrl := flag.String("imageUrl", "{PLACE_HOLDER}", "Url of image to use to decorate the document")

	// Parse the flag arguments from the terminal
	flag.Parse()

	// Validate template value
	if !(*template == "readme-template" || *template == "readme") {
		return userInput{}, errors.New("invalid template value. use readme-template or readme")
	}

	// Validate name
	if *name == "" {
		return userInput{}, errors.New("there was an error setting the document's name")
	}

	// Validate imageUrl
	if len(*imageUrl) == 0 {
		// TODO: add true image validation logic
		return userInput{}, errors.New("there was an error setting the image's Url")
	}

	// Return validated userInput
	return userInput{*template, *name, *imageUrl}, nil
}

func exitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
