package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
		_, err = generateAllFiles(inputs)
	} else {
		_, err = generateFile(inputs.template + ".md")
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

	imgPath := "https://github.com/WhitneyLampkin/devscriber/blob/main/assets/default-img.png?raw=true"

	// Define flags for the template, name and imageUrl arguments
	template := flag.String("template", "readme-template", "Template type to base the new document on")
	name := flag.String("name", "README", "Name of the new document")
	imageUrl := flag.String("imageUrl", imgPath, "Url of image to use to decorate the document")
	all := flag.Bool("all", false, "Generates all document template when true")

	// Parse the flag arguments from the terminal
	flag.Parse()

	// Validate template value
	/* if !(*template == "readme-template" || *template == "readme") {
		return userInput{}, errors.New("invalid template value. use readme-template or readme")
	} */

	// TODO: clean this up later
	if *template == "readme" {
		*template = "readme-template"
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
	return userInput{*template, *name, *imageUrl, *all}, nil
}

func generateAllFiles(inputs userInput) (bool, error) {
	// TODO: Remove this and simply return nil if there is no error...?
	isSuccessful := false

	// [OLD METHOD]
	// Set templates directory
	currentPath, err := os.Executable()
	check(err)
	wd, _ := filepath.Split(currentPath)
	tempPath := filepath.Join(wd, "templates")

	templateFiles, err := os.ReadDir(tempPath)
	check(err)

	for _, file := range templateFiles {
		_, err := generateFile(file.Name())
		check(err)
	}

	return isSuccessful, err
}

// Returns the filepath of the newly create file
func generateFile(templatePath string) (string, error) {
	isAvailable := false

	// Get the absolute path of the binary
	dir, err := filepath.Abs(filepath.Dir(templatePath))
	if err != nil {
		fmt.Println(err)
	}

	// Use the absolute path to reference the templates directory
	templatesDir := filepath.Join(dir, "templates")
	fmt.Println("Test: " + templatesDir)

	// TODO: Remove my failed attempt when I better understand the chosen solution.
	//currentPath, err := os.Executable()
	//check(err)
	//wd, _ := filepath.Split(currentPath)
	//tempPath := filepath.Join(wd, "templates")

	// Getting current templates to validate the templatePath
	availableTemplates, err := os.ReadDir("./templates")
	check(err)

	for _, file := range availableTemplates {
		if file.Name() == templatePath || strings.Contains(file.Name(), templatePath) {
			isAvailable = true
			break
		}
	}

	// TODO: Refactor later
	templatePath = "./templates/" + templatePath

	if !isAvailable {
		exitGracefully(fmt.Errorf("the %s filepath does not exist. please try again", templatePath))
	}

	// Check that template filepath still exists
	_, err = validateTemplatePath(templatePath)
	check(err)

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
	//newFilename := "README_" + timeString + ".md"
	var newFilename string

	switch templatePath {
	case "./templates/changelog-template.md":
		newFilename = "CHANGELOG_" + timeString + ".md"
	case "./templates/codeofconduct-template.md":
		newFilename = "CODE_OF_CONDUCT_" + timeString + ".md"
	case "./templates/contributing-template.md":
		newFilename = "CONTRIBUTING_" + timeString + ".md"
	case "./templates/readme-template.md":
		newFilename = "README_" + timeString + ".md"
	case "./templates/releasenotes-template.md":
		newFilename = "RELEASE_NOTES_" + timeString + ".md"
	}

	destination, err := os.Create(workingDir + "/" + newFilename)
	if err != nil {
		return "", errors.New("error creating the new file destination")
	}
	defer destination.Close()

	// Using io.Copy() since no replacements are being made at this time
	newFile, err := io.Copy(destination, newDocument)
	if newFile == 0 || err != nil {
		return "", errors.New("cannot create new file at this time, please try again")
	}

	fmt.Printf("\r\n\r\nSuccess: %s was added to the current directory\r\n", newFilename)

	return newFilename, nil
}

func storeTemplateFiles() error {
	templates, err := t.ReadDir("templates")
	check(err)

	for _, item := range templates {
		if item.IsDir() {
			continue
		}

		info, _ := item.Info()

		fmt.Println("Info: ", info)

		i, _ := t.Open("templates/" + item.Name())
		b := make([]byte, info.Size())
		i.Read(b)

		tFiles[item.Name()] = string(b)
	}

	return nil
}

// Ensures the chosen template still exists
func validateTemplatePath(templatePath string) (bool, error) {
	// Validate file exists
	if _, err := os.Stat(templatePath); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("file %s does not exist", templatePath)
	}

	// If no errors, the file is valid.
	return true, nil
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
