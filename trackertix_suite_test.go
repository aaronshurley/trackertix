package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTrackertix(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trackertix Suite")
}
