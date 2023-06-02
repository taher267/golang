package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "encoding/json"
	// "starconv"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init boos ver as a slice Book starut
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a new Book
func newBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	fmt.Println(book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Update Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

//	func handlerFunc(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprint(w, "<h1>Bismillahir Rohmanir Rohim</h1>")
//	}
var PORT string

func main() {
	// http.HandleFunc("/", handlerFunc)
	// fmt.Println("server is listerning on :3000")
	// http.ListenAndServe(":3000", nil)
	PORT = "3000"
	r := mux.NewRouter()
	// Mock Data @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "43984938", Title: "Book one", Author: &Author{Firstname: "Abu", Lastname: "Taher"}})
	books = append(books, Book{ID: "2", Isbn: "439849383", Title: "Book Two", Author: &Author{Firstname: "Nothing", Lastname: "Much"}})
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books", newBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	fmt.Println(`server is listerning on :` + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
