package service

import (
	"GoWithGin/entity"
	"GoWithGin/repository"
)

type MovieService interface {
	Save(entity.Movie) entity.Movie
	Update(entity.Movie) error
	Delete(entity.Movie) error
	FindById(Id int64) entity.Movie
	FindAll() []entity.Movie
}

type movieService struct {
	//movies []entity.Movie
	repository repository.MovieRepository
}

func NewMovieService(mr repository.MovieRepository) MovieService {
	return &movieService{
		repository: mr,
	}
}

func (service *movieService) Save(m entity.Movie) entity.Movie {
	//service.movies = append(service.movies, m)
	service.repository.Save(m)
	return m
}

func (service *movieService) FindAll() []entity.Movie {
	//return service.movies
	return service.repository.FindAll()
}
func (service *movieService) Update(video entity.Movie) error {
	service.repository.Update(video)
	return nil
}

func (service *movieService) Delete(video entity.Movie) error {
	service.repository.Delete(video)
	return nil
}
func (service *movieService) FindById(id int64) entity.Movie {
	return service.repository.FindById(id)
}
