package main

import (
	"log"
	"net/http"

	"github.com/Ritika-Agrawal811/learn-golang/project-book-management/internal/api/routes"
)

func main() {

	log.Println("Starting server at localhost:8080")

	router := routes.SetupRoutes()

	/* pass the custom router to ListenAndServe function */
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
