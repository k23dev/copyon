package main

import (
	"fmt"
	"os"

	"goshareit/pathscan"
	"goshareit/views"

	"github.com/labstack/echo/v4"
)

func main() {

	var path string
	var port string

	if len(os.Args) < 1 {
		path = "./"
	} else {
		path = os.Args[1]
	}

	// set port
	if len(os.Args) < 3 {
		port = "5686"
	} else {
		port = os.Args[2]
	}

	server := echo.New()

	server.GET("/", func(c echo.Context) error {
		ps := pathscan.New(path)
		view := views.Index(ps)

		SetupStaticRoutes(path, server)

		return Render(c, view)
	})

	// Start server
	server.Logger.Fatal(server.Start(":" + port))

}

func SetupStaticRoutes(path string, server *echo.Echo) {
	ps := pathscan.New(path)
	for _, dir := range ps.Directories {
		fmt.Println("/" + dir.Name)
		server.Static("/"+dir.Name, dir.Path)
	}
}
