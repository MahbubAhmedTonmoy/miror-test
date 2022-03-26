package controller

import (
	"GoWithGin/Cache"
	"GoWithGin/entity"
	"GoWithGin/service"
	"GoWithGin/validators"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type MovieController interface {
	FindAll() []entity.Movie
	FindById(ctx *gin.Context) entity.Movie
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controller struct {
	service service.MovieService
	cache   Cache.MovieCache
}

var validate *validator.Validate

func NewMovieController(s service.MovieService, c Cache.MovieCache) MovieController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: s,
		cache:   c,
	}
}

func (c *controller) FindAll() []entity.Movie {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) error {
	var movie entity.Movie
	err := ctx.ShouldBindJSON((&movie))
	if err != nil {
		return err
	}
	err = validate.Struct(movie)
	if err != nil {
		return err
	}

	c.service.Save(movie)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	ms := c.service.FindAll()
	data := gin.H{
		"title":  "Movie Page",
		"movies": ms,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func (c *controller) Update(ctx *gin.Context) error {
	var movie entity.Movie
	err := ctx.ShouldBindJSON(&movie)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	movie.ID = id
	err = validate.Struct(movie)
	if err != nil {
		return err
	}

	c.service.Update(movie)
	return nil
}
func (c *controller) Delete(ctx *gin.Context) error {
	var video entity.Movie
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id
	c.service.Delete(video)
	return nil
}

func (c *controller) FindById(ctx *gin.Context) entity.Movie {
	var video entity.Movie
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return entity.Movie{}
	}
	video.ID = id
	movie := c.cache.Get(strconv.Itoa(int(id)))
	if movie == nil {
		log.Println("from DB")
		result := c.service.FindById(int64(id))
		c.cache.Set(strconv.Itoa(int(id)), &result)

		return result
	} else {
		log.Println("from cache")
		result := *movie
		return result
	}
}
