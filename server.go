package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

type Book struct {
    Title string
    Author string
}

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
    }

    router := gin.Default()

    // Data API (TODO)
    api := router.Group("/v1")
    {
        api.GET("/books", func(c *gin.Context) {
            c.JSON(http.StatusOK, []Book{
                {"Gardens of the Moon", "Steven Erikson"},
                {"Longitude", "Dava Sobel"},
                {"Peter the Great", "Robert K. Massie"},
                {"The Two Koreas", "Don Oberdorfer"},
            })
        })
    }

    router.Run(":" + port)
}
