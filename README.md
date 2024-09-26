<h1 align="center" style="border-bottom: none">
    <a href="https://github.com/WhitneyLampkin/devscriber" target="_blank">
        <img alt="DevScriber" src="./assets/Designer.png" style="border-radius:50%; height:100px;">
    </a>
    <br>
    DevScriber
</h1>
<h3 align="center" style="border-bottom: none">
    A personal scribe for the lazy developer. :smirk:
</h3>

<hr />

<p align="center">
    Visit <a href="https://github.com/WhitneyLampkin/devscriber" target="_blank">DevScriber</a> for the full documentation, examples and guides.
</p>

<div align="center">

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
![GitHub language count](https://img.shields.io/github/languages/count/WhitneyLampkin/devscriber?label=Languages&color=yellow)
![GitHub contributors](https://img.shields.io/github/contributors/WhitneyLampkin/devscriber?label=Contributors&color=red)
![GitHub repo size](https://img.shields.io/github/repo-size/WhitneyLampkin/devscriber?label=Repo%20Size&color=teal)
![GitHub repo file count (file type)](https://img.shields.io/github/directory-file-count/WhitneyLampkin/devscriber?label=Files&color=purple)



</div>

DevScriber aims to provide a personal scribe to developers in the form of a command-line tool to simplify the documentation process.

For all intents and purposes, this is a fun, learning project. The goal is to build a consistent practice of daily Golang development. If development of the tool advances to the point of practical use, that would be an added bonus.

The planned features for DevScriber at the time of writing are as follows:

⭐ [README.md](./templates/readme-template.md) file generator<br />
⭐ [CHANGELOG.md](./templates/changelog-template.md) file generator<br />
⭐ [CONTRIBUTING.md](./templates/contributing-template.md) file generator<br />
⭐ [CODE_OF_CONDUCT.md](./templates/codeofconduct-template.md) file generator<br />
⭐ [RELEASE_NOTES.md](./templates/releasenotes-template.md) file generator<br />
⭐ GitHub user profile README.md file generator<br />
⭐ File updater (used to replace existing files)<br />
⭐ Customizable filenames, titles, images and shield badges

## Installation

The easiest way to install Devscriber is by downloading the appropriate pre-build binary for your system from the [GitHub release page](https://github.com/WhitneyLampkin/devscriber/releases).

## Usage

```txt
+-+-+-+-+-+-+-+-+-+-+
|D|e|v|S|c|r|i|b|e|r|
+-+-+-+-+-+-+-+-+-+-+

Usage: devscriber [OPTIONS]...

Options:
    -all
        Generates all available document templates when true (default: false)
    -imageUrl string
        Url of image to use to decorate the document (default: ./assests/default_image.png)
    -name string
        Name of the new document (default: README_[TIME_STRING].md")
    -template string
        Template type to base the new document on
        Available templates: changelog, codeofconduct, contributing, readme, releasenotes (default "readme")
```

## Testing

```shell
# Running Tests
ginkgo
# or
go test

# Running Code Coverage Scripts
$ go test -coverprofile=coverage.out ./...
$ go tool cover -func=coverage.out
$ go tool cover -html=coverage.out
```

## Fun Facts

This project is intended for personal use, but I still wanted to give it a meaningful name and custom icon to make it feel more _official_. There are a few hidden meanings in the tool's name:

- The tool's purpose is to make the creation of common development documentation easier, thus it is a scribe to developers or a dev's scribe.
- DevScribe was the first option but DevScriber had a better ring to it and still makes sense.
    - ...a scriber is a tool used to mark materials like wood, steel, metal or plastics. Similarly, documentation _marks_ the code that developer's write.
- DevScribe (DevScriber) is also a cute play on the word describe (describer).

Also, all design credit for the DevScriber logo and default image is due to the new [AI-powered Bing Search Engine](https://www.bing.com/search?q=Bing%20AI&showconv=1&form=MW00X7).

## Acknowledgements

Special thanks goes to [Ulises Martinez](https://github.com/mx-ulises), who provided the inspiration for this project. Thank you for acknowledging the effort I put into including READMEs with each of my GitHub repositories and inspiring me to make it easier for others to do the same.

## License

GNU General Public License 3.0, see [LICENSE](./LICENSE).

<h1 align="center" style="border-bottom: none; margin-top: 50px;">
    <a href="https://github.com/WhitneyLampkin/devscriber" target="_blank">
        <img alt="DevScriber" src="./assets/Designer.png" style="border-radius: 5%; height: 150px;">
    </a>
</h1>
