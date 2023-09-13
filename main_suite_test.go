package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDevscriber(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Devscriber Suite")
}
