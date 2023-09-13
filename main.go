package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// A userInput stores the options provided by the user.
type userInput struct {
	template string
	name     string
	imageUrl string
	// TODO: add optional user input variable for destination filepath
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

	// Generate new file
	generateFile(inputs)

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
	imageUrl := flag.String("imageUrl", "./assests/default_image.png", "Url of image to use to decorate the document")

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

// Returns the filepath of the newly create file
func generateFile(inputs userInput) (string, error) {
	templatePath := "./templates/readme-template.md"

	// TODO: Determine the template file to use with a switch statement when other templates are added
	// TODO: Conditionally reset the templatePath in the future
	if !(inputs.template == "readme-template" || inputs.template == "readme") {
		return "", errors.New("the template filepath is no longer valid. please try again")
	}

	// Check that template filepath still exists
	if _, err := validateTemplatePath(templatePath); err != nil {
		exitGracefully(err)
	}

	// Open template file
	newDocument, err := os.Open(templatePath)
	// Check for errors
	check(err)
	// Ensure the file is closed after usage
	defer newDocument.Close()

	// TODO: Create and call a helper function to replace templated values

	// Get working directory
	workingDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error setting %s as the working directory", workingDir)
	}

	// Create new destination file location
	currentTime := time.Now().UTC().Nanosecond()
	timeString := strconv.Itoa(currentTime)
	destination, err := os.Create(workingDir + "/" + "README_" + timeString + ".md")
	if err != nil {
		return "", errors.New("error creating the new file destination")
	}
	defer destination.Close()

	// Using io.Copy() since no replacements are being made at this time
	newFile, err := io.Copy(destination, newDocument)
	if newFile == 0 || err != nil {
		return "", errors.New("cannot create new file at this time, please try again")
	}

	fmt.Println(destination.Name())

	return "", nil
}

func validateTemplatePath(templatePath string) (bool, error) {
	// Validate file exists
	if _, err := os.Stat(templatePath); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("file %s does not exist", templatePath)
	}

	// If no errors, the file is valid.
	return true, nil
}

func exitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func check(e error) {
	if e != nil {
		exitGracefully(e)
	}
}
