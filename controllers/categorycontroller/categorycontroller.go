package categorycontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request){
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories" : categories,
	}

	tmpl, err := template.ParseFiles("views/categories/index.html")
	if err != nil{
		panic(err)
	}

	tmpl.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		tmpl, err := template.ParseFiles("views/categories/create.html")
		if err != nil{
			panic(err)
		}

		tmpl.Execute(w, nil)
	}

	if r.Method == "POST"{
		// Menampung hasil input User
		var category entities.Categories

		category.Name = r.FormValue("name")
		category.Created_at = time.Now()
		category.Updated_at = time.Now()

		report := categorymodel.Create(category)
		if !report {
			tmpl,_ := template.ParseFiles("views/category/add")
			tmpl.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}

	
}

func Edit(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		tmpl, err := template.ParseFiles("views/categories/edit.html")
		if err != nil{
			panic(err)
		}

		//Menangkap ID dari URL
		idString := r.URL.Query().Get("id")	
		id, err := strconv.Atoi(idString)
		if err != nil{
			panic(err)
		}

		category := categorymodel.Detail(id)
		data := map[string]any{
			"category" : category,
		}

		tmpl.Execute(w, data)
	}

	if r.Method == "POST"{
		var category entities.Categories

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil{
			panic(err)
		}

		category.Name = r.FormValue("name")
		category.Updated_at = time.Now()

		result := categorymodel.Update(id, category)
		if !result {
			http.Redirect(w, r, w.Header().Get("Refere"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request){
	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil{
		panic(err)
	}

	result := categorymodel.Delete(id)
	if result != nil{
		panic(result)
	}
	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}