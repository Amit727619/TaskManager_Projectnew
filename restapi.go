// main.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Book represents a simple book structure
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Go Programming", Author: "John Doe"},
	{ID: "2", Title: "Web Development", Author: "Jane Doe"},
}

// getAllBooks returns the list of all books
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)\
	
}

// getBookByID returns a book by ID
func getBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
		
	}

	// If no book is found, return an empty response
	json.NewEncoder(w).Encode(Book{})

}

func main() {
	// Define route handlers
	http.HandleFunc("/books", getAllBooks)
	http.HandleFunc("/book", getBookByID)

	// Start the server
	port := 8000
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
