package http

import (
	heavyapiload "req_parallel/heavyapiload"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc heavyapiload.UseCase) {
	h := NewHandler(uc)

	heavyApiLoadEndpoints := router.Group("/api/heavy_api_load")
	{
		heavyApiLoadEndpoints.GET("/test", h.Load)
	}
}
