/*
   Author: Maen Abu Hammour
   Date: October 5, 2023
   Description: This Go file, part of the "handlers" package, contains a set of HTTP request handlers for a backend service.
   		These handlers manage various book-related operations, including retrieving a list of all books, creating a new book, retrieving a book by its ID, updating a book by its ID, and deleting a book by its ID.
  		The handlers use the provided HTTP request and response objects to handle incoming requests, interact with the data storage system to perform CRUD (Create, Read, Update, Delete) operations on book data, and respond with JSON-encoded data or appropriate error messages.
 		Additionally, the code includes Swagger-style comments that document the handlers' functionality and expected input/output formats for API documentation purposes.
*/

package handlers

import (
	"book-info-service/src/model"
	"book-info-service/src/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	l4g "github.com/alecthomas/log4go"
)

// @Summary Get a list of all books
// @Description Get a list of all books.
// @Produce json
// @Success 200 {array} model.Book
// @Router /books [get]
func GetAllBooks(w http.ResponseWriter) {
	booksList, err := storage.DB.GetAllBooks()
	if err != nil {
		errMsg := fmt.Sprintf("%s", err)
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booksList)
}

// @Summary Create a new book
// @Description Create a new book.
// @Accept json
// @Param book body model.Book true "Book object to create"
// @Success 201 {object} model.Book
// @Router /books [post]
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook model.Book

	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	l4g.Debug("Book to be created: %s", newBook)
	id, err := storage.DB.CreateBook(newBook)
	if err != nil {
		errMsg := fmt.Sprintf("%s", err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	newBook.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

// @Summary Get a book by its ID
// @Description Get a book by its ID.
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} model.Book
// @Router /books/{id} [get]
func GetBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/books/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	book, err := storage.DB.GetBookByID(id)
	if err != nil {
		errMsg := fmt.Sprintf("%s", err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// @Summary Update a book by its ID
// @Description Update a book by its ID.
// @Accept json
// @Param id path int true "Book ID"
// @Param book body model.Book true "Updated book object"
// @Success 200 {object} model.Book
// @Router /books/{id} [put]
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/books/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var updatedBook model.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = storage.DB.UpdateBook(id, updatedBook)
	if err != nil {
		errMsg := fmt.Sprintf("%s", err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	updatedBook.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

// @Summary Delete a book by its ID
// @Description Delete a book by its ID.
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} model.Book
// @Router /books/{id} [delete]
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/books/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	book, err := storage.DB.DeleteBook(id)
	if err != nil {
		errMsg := fmt.Sprintf("%s", err)
		http.Error(w, errMsg, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
