// @title Example API
// @version 1.0
// @description Description of the example API
// @securityDefinitions.basic BasicAuth
package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	p "goserv/pkg/proxy"
	"time"
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

	router.GET("/starlite/hello", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/hello")
	})
	router.POST("/starlite/hello", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/hello")
	})

	router.GET("/starlite/address/geocoded", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/address/geocoded")
	})
	router.GET("/starlite/cart", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/cart")
	})
	router.GET("/starlite/cart/quote", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/cart/quote")
	})
	router.GET("/starlite/catalog/equipment", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/catalog/equipment")
	})
	router.GET("/starlite/catalog/equipment/classes", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/catalog/equipment/classes")
	})
	router.GET("/starlite/catalog/equipment/families", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/catalog/equipment/families")
	})
	router.GET("/starlite/catalog/equipment/categories", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/catalog/equipment/categories")
	})
	router.GET("/starlite/catalog/parts", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/catalog/equipment/parts")
	})
	router.POST("/starlite/pricebook", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/pricebook")
	})
	router.POST("/starlite/pricebook/realtime", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/pricebook/realtime")
	})
	router.GET("/starlite/client/token", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/client/token")
	})
	router.GET("/starlite/customer/contracts", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/customer/contracts")
	})
	router.GET("/starlite/customer/branches", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/customer/branches")
	})
	router.GET("/starlite/customer", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/customer")
	})
	router.GET("/starlite/doclib/assets", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/doclib/assets")
	})
	router.GET("/starlite/doclib/taxonomy", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/doclib/taxonomy")
	})
	router.GET("/starlite/customer/user/entitled", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/user/entitled")
	})
	router.GET("/starlite/inventory/availability", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/inventory/availability")
	})
	router.GET("/starlite/leads/dealers", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/leads/dealers")
	})
	router.GET("/starlite/leads/history", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/leads/history")
	})
	router.GET("/starlite/leads/statuses", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/leads/statuses")
	})
	router.GET("/starlite/mobile/search_nav", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/mobile/search_nav")
	})
	router.GET("/starlite/quote", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/quote")
	})
	router.DELETE("/starlite/quote", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/quote")
	})
	router.PUT("/starlite/quote", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/quote")
	})
	router.GET("/starlite/quote/search", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/quote/search")
	})
	router.GET("/starlite/quote/mine", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/quote/mine")
	})
	router.GET("/starlite/order/summaries", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/order/summaries")
	})
	router.GET("/starlite/partsfinder/serial", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/partsfinder/serial")
	})
	router.GET("/starlite/partsfinder/sku", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/partsfinder/sku")
	})
	router.GET("/starlite/partsfinder/unit", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/partsfinder/unit")
	})
	router.GET("/starlite/customer/skus", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/customer/skus")
	})
	router.GET("/starlite/user/permissions", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/user/permissions")
	})
	router.GET("/starlite/customer/user", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/customer/user")
	})
	router.POST("/starlite/ahrisys/match", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/arhisys/match")
	})
	router.PUT("/starlite/pns/token", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/pns/token")
	})
	router.POST("/starlite/ahrisys/search", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/ahrisys/search")
	})
	router.GET("/starlite/search", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/search")
	})
	router.POST("/starlite/pns/message", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/pns/message")
	})
	router.POST("/starlite/order", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/order")
	})
	router.POST("/starlite/order/validate", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/order/validate")
	})
	router.GET("/starlite/promos/validate", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/promos/validate")
	})
	router.GET("/starlite/apidocs/redoc", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/apidocs/redoc")
	})
	router.GET("/starlite/apidocs/openapi.json", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/apidocs/openapi.json")
	})
	router.GET("/starlite/apidocs/openapi.yaml", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/apidocs/openapi.yaml")
	})
	router.GET("/starlite/apidocs", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/apidocs")
	})
	router.GET("/starlite/apidocs/elements", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/apidocs/elements")
	})
	router.GET("/starlite/apidocs/swagger", func(c *gin.Context) {
		p.Proxy(c, "https://stageapi.daikincloud.io", "/starlite/1.0/apidocs/swagger")
	})

	router.Run("localhost:3000")
}
