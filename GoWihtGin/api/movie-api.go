package api

import (
	"GoWithGin/controller"
	"GoWithGin/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieApi struct {
	loginController controller.LoginController
	movieController controller.MovieController
}

func NewMovieAPI(lc controller.LoginController,
	mc controller.MovieController) *MovieApi {
	return &MovieApi{
		loginController: lc,
		movieController: mc,
	}
}

func (api *MovieApi) Authenticate(ctx *gin.Context) {
	token := api.loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, &dto.JWT{
			Token: token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, &dto.Response{
			Message: "Not Authorized",
		})
	}
}

func (api *MovieApi) GetMovies(ctx *gin.Context) {
	ctx.JSON(200, api.movieController.FindAll())
}
func (api *MovieApi) GetMovie(ctx *gin.Context) {
	ctx.JSON(200, api.movieController.FindById(ctx)) 
}
func (api *MovieApi) CreateMovie(ctx *gin.Context) {
	err := api.movieController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}

func (api *MovieApi) UpdateMovie(ctx *gin.Context) {
	err := api.movieController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}

func (api *MovieApi) DeleteMovie(ctx *gin.Context) {
	err := api.movieController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}
