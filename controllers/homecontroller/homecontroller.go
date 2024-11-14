package homecontroller

import (
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("views/home/index.html")
	if err != nil{
		panic(err)
	}

	tmpl.Execute(w, nil)
}