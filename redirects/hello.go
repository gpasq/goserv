package redirects

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloRedir(c *gin.Context) {
	j := `{"message":"Hello World","current_version":"Version 1.7.74","env":"stage"}`
	c.Header("Content-type", "application/xml")
	c.String(http.StatusOK, j)
}
