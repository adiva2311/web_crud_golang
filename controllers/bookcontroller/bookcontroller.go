package bookcontroller

import (
	"go-web-native/entities"
	"go-web-native/models/bookmodel"
	"go-web-native/models/categorymodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	books := bookmodel.GetAll()
	data := map[string]any{
		"books" : books,
	}

	tmpl, err := template.ParseFiles("views/books/index.html")
	if err != nil{
		panic(err)
	}

	tmpl.Execute(w, data)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil{
		panic(err)
	}

	book := bookmodel.Detail(id)
	data := map[string]any{
		"book" : book,
	}

	tmpl, err := template.ParseFiles("views/books/detail.html")
	if err != nil{
		panic(err)
	}

	tmpl.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		tmpl, err := template.ParseFiles("views/books/create.html")
		if err != nil{
			panic(err)
		}

		// Untuk Mengambil Semua Category
		categories := categorymodel.GetAll()
		data := map[string]any{
		"categories" : categories,
	}

		tmpl.Execute(w, data)
	}

	if r.Method == "POST"{
		var book entities.Books

		// Mengkonversi Tipe Data Category & Stock menjadi int, karena nilai yang dikirim dari Form berupa string
		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil{
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil{
			panic(err)
		}

		book.Name = r.FormValue("name")
		book.Category_id.Id = uint(categoryId)
		book.Stock = stock
		book.Description = r.FormValue("description")
		book.Created_at = time.Now()
		book.Updated_at = time.Now()

		report := bookmodel.Create(book)
		if !report {
			tmpl,_ := template.ParseFiles("views/books/add")
			tmpl.Execute(w, nil)
		}

		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		tmpl, err := template.ParseFiles("views/books/edit.html")
		if err != nil{
			panic(err)
		}

		//Menangkap ID dari URL
		idString := r.URL.Query().Get("id")	
		id, err := strconv.Atoi(idString)
		if err != nil{
			panic(err)
		}

		// Untuk Mengambil Semua Category
		categories := categorymodel.GetAll()
		book := bookmodel.Detail(id)
		data := map[string]any{
		"categories" : categories,
		"book" : book,
	}

		tmpl.Execute(w, data)
	}

	if r.Method == "POST"{
		var book entities.Books

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		// Mengkonversi Tipe Data Category & Stock menjadi int, karena nilai yang dikirim dari Form berupa string
		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil{
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil{
			panic(err)
		}

		book.Name = r.FormValue("name")
		book.Category_id.Id = uint(categoryId)
		book.Stock = stock
		book.Description = r.FormValue("description")
		book.Updated_at = time.Now()

		report := bookmodel.Update(id, book)
		if !report {
			tmpl,_ := template.ParseFiles("views/books/edit")
			tmpl.Execute(w, nil)
		}

		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := bookmodel.Delete(id); err != nil{
		panic(err)
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}