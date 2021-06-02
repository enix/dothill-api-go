package mock

import (
	"net/http"

	"github.com/enix/dothill-api-go/internal/mock/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	auth := controllers.NewAuthController()

	api := router.Group("api")
	{
		api.GET("/login/:hash", auth.Login)
	}

	router.Use(func(c *gin.Context) {
		if sessionKey, ok := c.Request.Header["Sessionkey"]; ok && len(sessionKey) > 0 {
			if _, ok := auth.Tokens[sessionKey[0]]; !ok {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	})

	router.NoRoute(func(c *gin.Context) {
		c.String(400, "")
	})

	return router
}
