package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(time.Now())
		}
	}()
	r := gin.Default()
	v1 := r.Group("/")
	v1.GET("ecard", func(c *gin.Context) {
		time.Sleep(time.Millisecond * 1500)
		c.String(http.StatusOK, "csb1")
	})

	r.Run(":11133")
}
