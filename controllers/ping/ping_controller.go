package ping
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping a
func Ping(c *gin.Context){
	c.String(http.StatusOK, "pong")
}