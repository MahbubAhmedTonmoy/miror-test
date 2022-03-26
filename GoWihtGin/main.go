package main

import (
	"GoWithGin/Cache"
	"GoWithGin/api"
	"GoWithGin/controller"

	//"GoWithGin/docs"
	"GoWithGin/middlewares"
	"GoWithGin/repository"
	"GoWithGin/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	//"github.com/pragmaticreviews/golang-gin-poc/docs" // Swagger generated files
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

var (
	movieRepo       repository.MovieRepository = repository.NewMovieRepository()
	movieService    service.MovieService       = service.NewMovieService(movieRepo)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()
	cache           Cache.MovieCache           = Cache.NewRedisCache("localhost:6379", 0, 10)
	loginController controller.LoginController = controller.NewLoginConroller(loginService, jwtService)
	movieController controller.MovieController = controller.NewMovieController(movieService, cache)
)

func logFile() {
	f, _ := os.Create("movie.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	logFile()
	//Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Pragmatic Reviews - Video API"
	docs.SwaggerInfo.Description = "Pragmatic Reviews - Youtube Video API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "pragmatic-video-app.herokuapp.com"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"https"}

	//defer videoRepository.CloseDB()
	server := gin.New()
	//server.Use(gin.Recovery(), middlewares.Logger(), middlewares.Auth(), gindump.Dump())
	server.Use(gin.Recovery(), middlewares.Logger())

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	// server.POST("/login", func(ctx *gin.Context) {
	// 	token := loginController.Login(ctx)
	// 	if token != "" {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"token": token,
	// 		})
	// 	} else {
	// 		ctx.JSON(http.StatusUnauthorized, nil)
	// 	}

	// })

	// apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	// {
	// 	apiRoutes.GET("/movies", func(ctx *gin.Context) {
	// 		ctx.JSON(200, movieController.FindAll())
	// 	})
	// 	apiRoutes.POST("/movies", func(ctx *gin.Context) {
	// 		err := movieController.Save(ctx)
	// 		if err != nil {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		} else {
	// 			ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is valid"})
	// 		}

	// 	})
	// 	apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
	// 		err := movieController.Update(ctx)
	// 		if err != nil {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		} else {
	// 			ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
	// 		}

	// 	})

	// 	apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
	// 		err := movieController.Delete(ctx)
	// 		if err != nil {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		} else {
	// 			ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
	// 		}

	// 	})

	// }
	movieAPI := api.NewMovieAPI(loginController, movieController)

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", movieAPI.Authenticate)
		}

		videos := apiRoutes.Group("/videos", middlewares.AuthorizeJWT())
		{
			videos.GET("", movieAPI.GetMovies)
			videos.GET(":id", movieAPI.GetMovie)
			videos.POST("", movieAPI.CreateMovie)
			videos.PUT(":id", movieAPI.UpdateMovie)
			videos.DELETE(":id", movieAPI.DeleteMovie)
		}
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/movies", movieController.ShowAll)
	}
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// We can setup this env variable from the EB console
	port := os.Getenv("PORT")

	// Elastic Beanstalk forwards requests to port 5000
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)

}
