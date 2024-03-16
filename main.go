package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//go:embed assets
//var a embed.FS

//go:embed templates
var t embed.FS

// Store template files
var tFiles = make(map[string]string)

// A userInput stores the options provided by the user
type userInput struct {
	template string
	name     string
	imageUrl string
	all      bool
}

// The entry point of the application
func main() {
	// Help info displayed when the --help flag is provided
	flag.Usage = func() {
		fmt.Println("+-+-+-+-+-+-+-+-+-+-+")
		fmt.Println("|D|e|v|S|c|r|i|b|e|r|")
		fmt.Println("+-+-+-+-+-+-+-+-+-+-+")
		fmt.Printf("\r\nUsage: %s [OPTIONS]...\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Get user input
	inputs, err := getUserInput()
	check(err)

	fmt.Println("+-+-+-+-+-+-+-+-+-+-+")
	fmt.Println("|D|e|v|S|c|r|i|b|e|r|")
	fmt.Println("+-+-+-+-+-+-+-+-+-+-+")

	fmt.Println("\r\nGenerating new file with the following values:")
	fmt.Printf("- name: %s", inputs.name)
	fmt.Printf("\r\n- template: %s", inputs.template)
	fmt.Printf("\r\n- imageUrl: %s", inputs.imageUrl)

	err = storeTemplateFiles()
	check(err)

	// Generate new file(s)
	if inputs.all {
		_, err = generateAllFiles()
	} else {
		_, err = generateFile(inputs.template)
	}
	check(err)
}

// Gets user input from the terminal and validates it
// Returns a validated userInput struct
func getUserInput() (userInput, error) {
	// Validate the number of required arguments
	if len(os.Args) == 0 {
		return userInput{}, errors.New("the template argument is required")
	}

	// <TODO> Use go:embed to allow users to choose from different images
	imgPath := "https://github.com/WhitneyLampkin/devscriber/blob/main/assets/default-img.png?raw=true"

	// Define flags for the template, name and imageUrl arguments
	template := flag.String("template", "readme", "Template type to base the new document on")
	name := flag.String("name", "README", "Name of the new document")
	imageUrl := flag.String("imageUrl", imgPath, "Url of image to use to decorate the document")
	all := flag.Bool("all", false, "Generates all document template when true")

	// Parse the flag arguments from the terminal
	flag.Parse()

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
	return userInput{*template, *name, *imageUrl, *all}, nil
}

func generateAllFiles() (bool, error) {
	// TODO: Remove this and simply return nil if there is no error...?
	isSuccessful := false

	for n := range tFiles {
		_, err := generateFile(n)
		fmt.Println()
		check(err)
	}

	return isSuccessful, nil
}

// Returns the filepath of the newly create file
func generateFile(templatePath string) (string, error) {
	isAvailable := false

	for name := range tFiles {
		if name == templatePath || strings.Contains(name, templatePath) {
			templatePath = name
			isAvailable = true
			break
		}
	}

	if !isAvailable {
		exitGracefully(fmt.Errorf("the %s filepath does not exist. please try again", templatePath))
	}

	// Get working directory
	workingDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error setting %s as the working directory", workingDir)
	}

	// Create new destination file location
	currentTime := time.Now().UTC().Nanosecond()
	timeString := strconv.Itoa(currentTime)

	var newFilename string

	switch templatePath {
	case "changelog-template.md":
		newFilename = "CHANGELOG_" + timeString + ".md"
	case "codeofconduct-template.md":
		newFilename = "CODE_OF_CONDUCT_" + timeString + ".md"
	case "contributing-template.md":
		newFilename = "CONTRIBUTING_" + timeString + ".md"
	case "readme-template.md":
		newFilename = "README_" + timeString + ".md"
	case "releasenotes-template.md":
		newFilename = "RELEASE_NOTES_" + timeString + ".md"
	}

	fmt.Println(workingDir)
	fmt.Println(newFilename)
	fmt.Println(templatePath)
	destination, err := os.Create(workingDir + "/" + newFilename)
	if err != nil {
		return "", errors.New("error creating new file destination")
	}
	defer destination.Close()

	destination.WriteString(tFiles[templatePath])

	// Using io.Copy() since no replacements are being made at this time
	// newFile, err := io.Copy(destination, newDocument)
	// if newFile == 0 || err != nil {
	// 	return "", errors.New("cannot create new file at this time, please try again")
	// }

	fmt.Printf("\r\n\r\nSuccess: %s was added to the current directory\r\n", newFilename)

	return newFilename, err
}

func storeTemplateFiles() error {
	templates, err := t.ReadDir("templates")
	check(err)

	for _, item := range templates {
		if item.IsDir() {
			continue
		}

		info, _ := item.Info()

		// Use for Debugging
		// fmt.Println("Info: ", info)

		i, _ := t.Open("templates/" + item.Name())
		b := make([]byte, info.Size())
		i.Read(b)

		tFiles[item.Name()] = string(b)
	}

	return nil
}

// Terminates the execution of the program when unexpected errors occur
func exitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

// Checks for errors and terminates the application
func check(e error) {
	if e != nil {
		exitGracefully(e)
	}
}
