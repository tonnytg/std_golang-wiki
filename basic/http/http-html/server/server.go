package server

import (
	"html/template"
	"log"
	"net/http"
)

type Info struct {
	ID      int
	Title   string
	Content string
}

var tmpl *template.Template

func FormHandler(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Method)

	if r.Method == "POST" {
		// get user and password form query url
		r.ParseForm()
		user := r.FormValue("user")
		pass := r.FormValue("password")

		if user == "admin" && pass == "admin" {
			log.Println("Login Success")
		} else {
			log.Println("Login Failed")
		}
	}
	// use template for form.html
	tmpl := template.Must(template.New("form.html").ParseFiles("./server/templates/form.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func PersonHandler(w http.ResponseWriter, req *http.Request) {
	news := []Info{
		{1, "Vaga", "Test"},
		{2, "Vaga2", "Test2"},
		{3, "Vaga3", "Test3"},
	}

	tmpl := template.Must(template.New("home.html").ParseFiles("./server/templates/home.html"))
	if err := tmpl.Execute(w, news); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WebServer() {

	http.HandleFunc("/", PersonHandler)
	http.HandleFunc("/form", FormHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
