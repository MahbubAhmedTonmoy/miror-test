package test

import (
	"GoWithGin/entity"
	"GoWithGin/repository"
	"GoWithGin/service"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	TITLE       = "Video Title"
	DESCRIPTION = "Video Description"
	URL         = "https://youtu.be1/JgW-i2QjgHQ"
	FIRSTNAME   = "John"
	LASTNAME    = "Doe"
	EMAIL       = "jdoe1@mail.com"
)

var testMovie entity.Movie = entity.Movie{
	Title:       TITLE,
	Description: DESCRIPTION,
	URL:         URL,
	Author: entity.Person{
		FirstName: FIRSTNAME,
		LastName:  LASTNAME,
		Email:     EMAIL,
	},
}

var _ = Describe("Movie Service", func() {
	var (
		mr repository.MovieRepository
		ms service.MovieService
	)

	BeforeSuite(func() {
		mr = repository.NewMovieRepository()
		ms = service.NewMovieService(mr)
	})
	Describe("Fetching all existing movies", func() {
		Context("if there is a movie in DB", func() {

			BeforeEach(func() {
				ms.Save(testMovie)
			})

			It("should return at least one element", func() {
				movielist := ms.FindAll()
				Ω(movielist).ShouldNot(BeEmpty())
			})

			It("should map the fields correctly", func() {
				firstVideo := ms.FindAll()[0]

				Ω(firstVideo.Title).Should(Equal(TITLE))
				Ω(firstVideo.Description).Should(Equal(DESCRIPTION))
				Ω(firstVideo.URL).Should(Equal(URL))
				Ω(firstVideo.Author.FirstName).Should(Equal(FIRSTNAME))
				Ω(firstVideo.Author.LastName).Should(Equal(LASTNAME))
				Ω(firstVideo.Author.Email).Should(Equal(EMAIL))
			})

			AfterEach(func() {
				video := ms.FindAll()[0]
				ms.Delete(video)
			})

		})

		Context("if there are no movie in DB", func() {
			It("", func() {
				movies := ms.FindAll()
				Ω(movies).Should(BeEmpty())
			})
		})
	})
})
