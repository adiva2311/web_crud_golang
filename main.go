package main

import (
	"go-web-native/config"
	"go-web-native/controllers/bookcontroller"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. Home Page
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)
	
	// 3. Books
	http.HandleFunc("/books", bookcontroller.Index)
	http.HandleFunc("/books/detail", bookcontroller.Detail)
	http.HandleFunc("/books/add", bookcontroller.Add)
	http.HandleFunc("/books/edit", bookcontroller.Edit)
	http.HandleFunc("/books/delete", bookcontroller.Delete)

	log.Println("Server is Running on Port 9000")
	http.ListenAndServe("127.0.0.1:9000", nil)
}