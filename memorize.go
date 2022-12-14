package main

import(
	"memorize/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	return port
}

func main() {
	router := gin.Default()
	// router.LoadHTMLGlob("templates/**/*.tmpl")
	controllers.Setup(router)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Nothing to see here.")
	})

	err := router.Run(":" + getPort())
	if err != nil {
		panic(err)
	}
}


// vi:syntax=go

