package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Book Struct
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`

}

//Author struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

//Init books variable as a slice Book struct
var books []Book

func main() {
	//Init Router
	r := mux.NewRouter()

	//Mock Data - @todo Implement DB
	books = append(books, Book{ID: "1", Isbn: "123213", Title: "One",
	 Author: &Author{FirstName: "first", LastName: "last"}})

	 books = append(books, Book{ID: "2", Isbn: "21323", Title: "Book Two", 
	 Author: &Author{FirstName:"firstSec", LastName: "LastSec"  }})

	 books = append(books, Book{ID: "3", Isbn: "22323", Title: "Book Three", 
	 Author: &Author{FirstName:"firstThird", LastName: "LastThird"  }})

	 books = append(books, Book{ID: "4", Isbn: "32343", Title: "Book Four", 
	 Author: &Author{FirstName:"firstFour", LastName: "LastFour"  }})

	 books = append(books, Book{ID: "5", Isbn: "546546", Title: "Book Five", 
	 Author: &Author{FirstName:"firstFive", LastName: "LastFive"  }})

	//Router Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}