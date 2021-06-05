package main

import (
    "fmt"

    "example.com/greetings"
    "github.com/gin-gonic/gin"
    "rsc.io/quote"
)

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)

    r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}