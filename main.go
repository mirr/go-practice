package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Bookmark Struct (Model)
type Bookmark struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `jsong:"url"`
	User        *User  `json:"user"`
}

// User Struct (Model)
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// RequestParams Struct
type RequestParams struct {
	ID int `json:"id"`
}

// Collection of Bookmarks
var bookmarks []Bookmark

// Get all Bookmarks
func getBookmarks(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Header().Set("X-Total", "2")

	// Response
	json.NewEncoder(responseWriter).Encode(bookmarks)
}

// Get a Bookmark
func getBookmark(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	// Get request as params
	params := mux.Vars(request)

	// Convert Param ID to int
	paramID, err := strconv.Atoi(params["id"])

	// Check if not null
	if err == nil {
		// Loop through bookmarks and find with id
		for _, item := range bookmarks {
			if item.ID == paramID {
				json.NewEncoder(responseWriter).Encode(item)
				return
			}
		}

		// Response
		json.NewEncoder(responseWriter).Encode(params)
	}
}

// Create a Bookmark
func createBookamrk(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	var bookmark Bookmark
	_ = json.NewDecoder(request.Body).Decode(&bookmark)

	// Generate fake bookmark ID
	bookmark.ID = rand.Intn(1000000000)

	// Append to bookmarks
	bookmarks = append(bookmarks, bookmark)

	// Response
	json.NewEncoder(responseWriter).Encode(bookmark)
}

// Update a Bookmark
func updateBookmark(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	paramID, err := strconv.Atoi(params["id"])

	if err == nil {
		for index, item := range bookmarks {
			if item.ID == paramID {
				bookmarks = append(bookmarks[:index], bookmarks[index+1:]...)
				var bookmark Bookmark
				_ = json.NewDecoder(request.Body).Decode(&bookmark)

				// Generate fake bookmark ID
				bookmark.ID = paramID

				// Append to bookmarks
				bookmarks = append(bookmarks, bookmark)

				// Response
				json.NewEncoder(responseWriter).Encode(bookmark)
				return
			}
		}

		// Response
		json.NewEncoder(responseWriter).Encode(bookmarks)
	}

}

// Delete a Bookmark
func deleteBookmark(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	paramID, err := strconv.Atoi(params["id"])

	if err == nil {
		for index, item := range bookmarks {
			if item.ID == paramID {
				bookmarks = append(bookmarks[:index], bookmarks[index+1:]...)
				break
			}
		}

		// Response
		json.NewEncoder(responseWriter).Encode(bookmarks)
	}
}

// Main
func main() {
	// Init the Router
	router := mux.NewRouter()

	// Mock Data
	bookmarks = append(bookmarks, Bookmark{
		ID:          1,
		Title:       "Go Lang",
		Description: "The Go Programming Language",
		URL:         "https://golang.org/",
		User: &User{
			Name:  "Mirr",
			Email: "Mirr@test.test",
		},
	})
	bookmarks = append(bookmarks, Bookmark{
		ID:          2,
		Title:       "NodeJS",
		Description: "The NodeJS Programming Language",
		URL:         "https://nodejs.org/",
		User: &User{
			Name:  "Mirr",
			Email: "Mirr@test.test",
		},
	})

	// Route handlers / Endpoint
	router.HandleFunc("/api/bookmarks", getBookmarks).Methods("GET")
	router.HandleFunc("/api/bookmarks/{id}", getBookmark).Methods("GET")
	router.HandleFunc("/api/bookmarks", createBookamrk).Methods("POST")
	router.HandleFunc("/api/bookmarks/{id}", updateBookmark).Methods("PUT")
	router.HandleFunc("/api/bookmarks/{id}", deleteBookmark).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
