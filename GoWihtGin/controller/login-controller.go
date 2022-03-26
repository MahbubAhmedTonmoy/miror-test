package controller

import (
	"GoWithGin/dto"
	"GoWithGin/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewLoginConroller(ls service.LoginService,
	jwtS service.JWTService) LoginController {
	return &loginController{
		loginService: ls,
		jwtService:   jwtS,
	}
}

func (c *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := c.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return c.jwtService.GenerateToken(credentials.Username, true)
	}
	return "not authorize"
}
