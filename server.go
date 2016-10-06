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

    // Main page (not on "/" as Static can't be overridden for the data API)
    router := gin.Default()
    router.Static("/main", "./dist")
    router.GET("/", func(c *gin.Context) {
        c.Redirect(http.StatusSeeOther, "/main")
    })

    // Data API (TODO)
    api := router.Group("/data")
    {
        api.GET("/test", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "message": "pong",
            })
        })
    }

    router.Run(":" + port)
}
