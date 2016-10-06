package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
    }

    router := gin.Default()

    // Data API (TODO)
    api := router.Group("/data")
    {
        api.GET("/test", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "message": "OK!",
            })
        })
    }

    router.Run(":" + port)
}
