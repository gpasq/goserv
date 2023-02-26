// @title Example API
// @version 1.0
// @description Description of the example API
// @securityDefinitions.basic BasicAuth
package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/hello", helloRedir)
	router.GET("/hello", func(c *gin.Context) {
		j := `{"message":"Hello World","current_version":"Version 1.7.74","env":"stage"}`
		c.Header("Content-type", "application/xml")
		c.String(http.StatusOK, j)
	})

	router.Run("localhost:3000")
}
