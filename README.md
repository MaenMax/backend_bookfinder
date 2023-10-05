# Backend Developer Project: Book Information Service

This project create a backend service that provides CRUD operations for book information. The project creates two sets of APIs:

RESTful APIs:
----
-  POST      ```/books```
-  PUT       ```/books/{id}```
-  GET       ```/books/```
-  GET       ```/books/{id}```
-  DELETE    ```/books/{id}```


GraphQL APIs:
----
- POST    ```/books-graphql```

Those two APIs will run on localhost port ```8080```


### Table of Contents

1. [Installing Golang](#installing-golang)
2. [Building the project](#building-project)
3. [Running the system](#running-the-system)
4. [API Documentation](#api-documentation)
5. [Examples](#examples)

### Installing Golang
Before building/running the server, you will need to install Golang on your machine. For more information on how to install it on Linux/Mac/Windows, please visit: https://go.dev/doc/install


### Building Project
To build the project (generate the Microservice binary), go the home directory of the repo/projct and run:

```bash
$ go mod vendor
```

```bash
$ make
```
To generate a tar.bz2 file for deployment, run:
```bash
$ make deploy
```
To create a docker image, run:
```bash
$ make docker
```
To delete the generated file (binaries, docker, tar.bz2):
```bash
$ make clean
```


### Running the system

The generated binary will be stored in $HOME/bin/book_finder, where $HOME is the home directory of the repo/project

To run it from command line (Mac and Linux):
```bash
$ ./bin/book_finder
```
The process will start and the following logs will show:

```
GraphQL server is running on :8080/books-graphql
RESTful server is running on :8080/books
```



### API Documentation

API documentation can be found in $home/swagger.yaml, where $HOME is the home directory of the repo/project



### Examples

For a server running on localhost:8080

#### RESTful requests

##### Create a new book
```
POST http://localhost:8080/books
{
  "title": "Book title",
  "author": "Author",
  "date_published": "2023-10-15",
  "book_cover_url": "https://example.com/new-book-cover.jpg"
}
```

##### Update existing book
```
PUT http://localhost:8080/books/1
{
  "title": "Updated Title",
  "author": "Author updated",
  "date_published": "2023-11-01",
  "book_cover_url": "https://example.com/updated-cover.jpg"
}
```
##### Delete Book
```
DELETE http://localhost:8080/books/1
{
  "title": "Book title",
  "author": "Author",
  "date_published": "2023-10-15",
  "book_cover_url": "https://example.com/new-book-cover.jpg"
}
```
##### Get Book by its ID
```
GET http://localhost:8080/books/1
```
##### Get All Books
```
GET http://localhost:8080/books
```


#### GraphQL requests

##### Create a new book
```
POST http://localhost:8080//books-graphql
{
  "query": "mutation($title: String!, $author: String!, $datePublished: String!, $bookCoverURL: String!) { createBook(Title: $title, Author: $author, DatePublished: $datePublished, BookCoverURL: $bookCoverURL) { ID, Title, Author, DatePublished, BookCoverURL } }",
  "variables": {
    "title": "Title",
    "author": "Maen Abu Hammour",
    "datePublished": "09/15/2023",
    "bookCoverURL": "https://example.com/book-cover.jpg"
  }
}
```

##### Update existing book
```
POST http://localhost:8080/books-graphql
{
  "query": "mutation ($ID: Int!, $Title: String, $Author: String, $DatePublished: String, $BookCoverURL: String) { updateBook(ID: $ID, Title: $Title, Author: $Author, DatePublished: $DatePublished, BookCoverURL: $BookCoverURL) { ID, Title, Author, DatePublished, BookCoverURL } }",
  "variables": {
    "ID": 1,  // Replace with the actual ID of the book you want to update
    "Title": "Updated Title",
    "Author": "Updated Author",
    "DatePublished": "2023-10-01",
    "BookCoverURL": "https://example.com/updated-cover.jpg"
  }
}
```
##### Delete Book
```
POST http://localhost:8080/books-graphql
{
  "query": "mutation ($ID: Int!) { deleteBook(ID: $ID) { ID, Title, Author, DatePublished, BookCoverURL } }",
  "variables": {
    "ID": 1  // Replace with the actual ID of the book you want to delete
  }
}
```
##### Get Book by its ID
```
POST http://localhost:8080/books-graphql
{
  "query": "query ($ID: Int!) { book(id: $ID) { ID, Title, Author, DatePublished, BookCoverURL } }",
  "variables": {
    "ID": 1
  }
}
```
##### Get All Books
```
POST http://localhost:8080/books-graphql
```
