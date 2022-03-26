package handler

import (
	"log"
	"net/http"
	"restapigogin/model"
	"restapigogin/model/apperrors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Me(c *gin.Context) {

	user, exists := c.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})

		return
	}
	id := user.(*model.User).ID
	ctx := c.Request.Context()
	u, err := h.userService.Get(ctx, id)

	if err != nil {
		log.Printf("Unable to find user: %v\n%v", id, err)
		e := apperrors.NewNotFound("user", id.String())

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
