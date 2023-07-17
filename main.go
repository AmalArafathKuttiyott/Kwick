package main

import (
	"fmt"
	"kwick/helper/initializer"
	"kwick/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnv()
	initializer.MigrateTables()
}
func main() {
	r := gin.Default()
	routes.Routes(r)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	err := r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
