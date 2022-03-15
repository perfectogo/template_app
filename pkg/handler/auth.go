package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	template "github.com/perfectogo/template_app"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input template.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var input template.User
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
}
