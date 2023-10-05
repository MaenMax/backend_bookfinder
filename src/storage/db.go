/*
   Author: Maen Abu Hammour
   Date: October 5, 2023
   Description: Manages book data storage and retrieval for a backend service. It defines a "Database" interface for various database operations and provides an in-memory database ("InMemoryDB") implementation as an implementation of that interface just for the purpose of this excercise. It provides a layer of abstraction, where another DBs (like SQL and others) can be added. The code initializes and offers functions for CRUD (Create, Read, Update, Delete) operations on books. Additionally, it includes a helper function for generating unique book IDs. This file serves as the data storage backbone for the application.
*/

package storage

import (
	"book-info-service/src/model"
	"errors"

	l4g "github.com/alecthomas/log4go"
)

var (
	nextBookID = 1
	DB         Database
)

func init() {

	// Initialize storage with an instance of InMemoryDB
	DB = &InMemoryDB{}

	// Initialize Database
	if err := DB.Initialize(); err != nil {
		l4g.Exitf("Failed to initialize database: %v", err)
	}
}

// Define an interface for the database operations. Using this method, the design will be able to support various types of databases
type Database interface {
	Initialize() error
	GetAllBooks() ([]model.Book, error)
	GetBookByID(id int) (model.Book, error)
	CreateBook(book model.Book) (int, error)
	UpdateBook(id int, book model.Book) error
	DeleteBook(id int) (model.Book, error)
}

// This is an in-memory db, specifically for this excercise
type InMemoryDB struct {
	db map[int]model.Book
}

/** In the same way above, you may want to define another DB which implement the "Database" interface (in real case scenario)
For example:

type SQLDatabase struct {
    db *sql.DB
}
**/

func (I *InMemoryDB) Initialize() error {
	I.db = make(map[int]model.Book)
	return nil
}

func (I *InMemoryDB) GetAllBooks() ([]model.Book, error) {
	booksList := make([]model.Book, 0, len(I.db))
	for _, book := range I.db {
		booksList = append(booksList, book)
	}
	return booksList, nil
}

func (I *InMemoryDB) GetBookByID(id int) (model.Book, error) {
	l4g.Debug("Getting book: %v", id)
	book, found := I.db[id]
	if !found {
		return model.Book{}, errors.New("Book not found")
	}
	book.ID = id
	return book, nil
}

func (I *InMemoryDB) CreateBook(book model.Book) (int, error) {
	l4g.Debug("CreateBook: %s", book)
	book.ID = GetNextBookID()
	I.db[book.ID] = book
	return book.ID, nil
}

func (I *InMemoryDB) UpdateBook(id int, book model.Book) error {
	l4g.Debug("Updating book with ID: %v", id)
	_, err := I.GetBookByID(id)
	if err != nil {
		l4g.Debug("Book with ID: %v was not found")
		return err
	}
	I.db[id] = book
	return nil
}

func (I *InMemoryDB) DeleteBook(id int) (model.Book, error) {
	book, err := I.GetBookByID(id)
	if err != nil {
		l4g.Debug("Book with ID: %v was not found")
		return model.Book{}, err
	}

	delete(I.db, id)
	return book, nil
}

// A helper function
func GetNextBookID() int {
	nextID := nextBookID
	nextBookID++
	return nextID
}
