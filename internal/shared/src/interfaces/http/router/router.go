package router

import (
	"go-wallet/internal/shared/src/interfaces/http/dependencies"
	"go-wallet/internal/shared/src/interfaces/http/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(d *dependencies.Dependencies) *gin.Engine {

	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Set(&dependencies.Dependencies{}, d)
		ctx.Next()
	})
	r.Use(middleware.SetRequestContext)
	r.Use(middleware.HandleError)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message": "pong"})
	})

	return r
}
