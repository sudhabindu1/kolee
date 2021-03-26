package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"NoOfDays": getDays(),
		})
	})
	router.Run(":8080")
}

func getDays() int {
	now := time.Now()

	format := "2006-01-02 15:04:05"
	then, _ := time.Parse(format, "2019-11-22 00:00:00")

	diff := now.Sub(then)

	//func Since(t Time) Duration
	//Since returns the time elapsed since t.
	//It is shorthand for time.Now().Sub(t).

	return int(diff.Hours() / 24)
}
