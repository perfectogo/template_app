package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateList(ctx *gin.Context) {
	id, _ := ctx.Get(userCtx)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAllLists(ctx *gin.Context) {

}

func (h *Handler) GetListById(ctx *gin.Context) {

}

func (h *Handler) UpdateList(ctx *gin.Context) {

}

func (h *Handler) DeleteList(ctx *gin.Context) {

}
