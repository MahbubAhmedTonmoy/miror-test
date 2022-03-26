package test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMovieService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Video Service Test Suite")
}
