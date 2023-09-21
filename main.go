package main

import (
	"todo/cmd"
	_ "todo/docs"
)

// @title						to-do-list
// @version					1.0
// @description				Swagger страница
// @BasePath					/
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	cmd.Run()
}
