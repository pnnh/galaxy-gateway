package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const IndexPageSize = 10

type indexHandler struct {
}

func (s *indexHandler) Query(gctx *gin.Context) {
	gctx.JSON(http.StatusOK, gin.H{"message": "galaxy-gateway"})
}

func NewIndexHandler() *indexHandler {
	return &indexHandler{}
}
