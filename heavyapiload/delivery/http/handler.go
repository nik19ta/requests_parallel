package http

import (
	"net/http"
	heavyapiload "req_parallel/heavyapiload"

	gin "github.com/gin-gonic/gin"
)

type Handler struct {
	useCase heavyapiload.UseCase
}

func NewHandler(useCase heavyapiload.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Load(c *gin.Context) {

	req, err := h.useCase.Load()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Error"})
	}

	c.JSON(http.StatusOK, gin.H{"msg": req})
}
