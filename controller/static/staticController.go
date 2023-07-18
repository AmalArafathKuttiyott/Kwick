package staticController

import (
	"path"

	"github.com/gin-gonic/gin"
)

func HandleStaticFiles(ctx *gin.Context) {
	// Get the path to the static folder relative to the current directory
	dir := "public/images/"

	// Strip the "/static/" prefix from the URL path and create a new request URL with the modified path
	filepath := ctx.Param("filename")

	fullPath := path.Join(dir, filepath)

	// Serve the static files
	ctx.File(fullPath)
}
