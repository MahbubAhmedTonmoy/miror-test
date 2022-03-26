package handler

import (
	"restapigogin/model"

	"github.com/gin-gonic/gin"
)

//Handler struct holds required services for handler to function
type Handler struct {
	userService model.UserService
}

//config will hold services that will eventually be injected into this
//handler layer on handler initialization
type Config struct {
	R           *gin.Engine
	userService model.UserService
}

func NewHandler(c *Config) {
	h := &Handler{
		userService: c.userService,
	}

	g := c.R.Group("/api/account")

	g.GET("/me", h.Me)
}
