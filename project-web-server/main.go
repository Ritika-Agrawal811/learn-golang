package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	// password := r.FormValue("password")

	if name == "" || email == "" {
		return
	}

	log.Printf("Received form submission: Name=%s, Email=%s", name, email)

	fmt.Fprintf(w, "POST request successful!\n")
	fmt.Fprintf(w, "Received users details ->\n")
	fmt.Fprintf(w, "Name: %s \n", name)
	fmt.Fprintf(w, "Email: %s \n", email)

}

func main() {
	// Serve static files
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Register routes
	http.HandleFunc("/register", formHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
