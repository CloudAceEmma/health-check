package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-ping/ping"
)

var domain string = "8.8.8.8"

func main() {
	r := gin.Default()
	r.GET("/health-check", connectivityCheck)
	r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:80
}

func connectivityCheck(c *gin.Context) {
	pinger, err := ping.NewPinger(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	pinger.Count = 1
	pinger.Timeout = 3000000000
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	stats := pinger.Statistics()
	if stats.PacketLoss == 100 {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, pinger.IPAddr())
}
