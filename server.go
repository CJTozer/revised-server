package main

import (
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/itsjamie/gin-cors"
)

type Book struct {
    Title string `json:"title"`
    Author string `json:"author"`
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
        // API allows cross-origin requests
        api.Use(cors.Middleware(cors.Config{
            Origins:        "*",
            Methods:        "GET, PUT, POST, DELETE",
            RequestHeaders: "Origin, Authorization, Content-Type",
            ExposedHeaders: "",
            MaxAge: 50 * time.Second,
            Credentials: true,
            ValidateHeaders: false,
        }))

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
