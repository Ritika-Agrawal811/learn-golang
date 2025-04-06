package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/learn-golang/project-movies-crud/api"
)

// @title Movie API
// @version 1.0
// @description This is a sample server for movies.
// @host localhost:8000
// @BasePath /api/v1
func main() {
	fmt.Println("starting server at port: 8000")

	router := api.SetupRoutes()

	/* pass the custom router to http.ListenAndServe method */
	if err := http.ListenAndServe(":8000", router); err != nil {
		fmt.Println("failed")
		log.Fatal(err)
	}

}
