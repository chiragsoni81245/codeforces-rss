package server

import (
	"net/http"

	"github.com/chiragsoni81245/codeforces-rss/internal/config"
	"github.com/chiragsoni81245/codeforces-rss/internal/feed"
	"github.com/gin-gonic/gin"
)

// StartServer initializes the HTTP server
func StartServer() {
    gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
    config := config.LoadConfig()

	// RSS Endpoint
	router.GET("/rss", func(c *gin.Context) {
		rss, err := feed.GenerateRSS(config.Tags, config.ExcludedTags, config.MinRating, config.MaxRating)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Header("Content-Type", "application/rss+xml")
		c.String(http.StatusOK, rss)
	})

	router.Run(":8080")
}

