package main

import (
	"net/http"
	"fmt"
	"time"
	"html/template"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

//Go application entrypoint
func main() {

	welcome := Welcome{"Viha", time.Now().Format(time.Stamp)}


	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))


	http.HandleFunc("/welcome" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/welcome.html"))

		if err := templates.ExecuteTemplate(w, "welcome.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})


	http.HandleFunc("/login" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/login.html"))

		if err := templates.ExecuteTemplate(w, "login.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	


	http.HandleFunc("/verify-login" , func(w http.ResponseWriter, r *http.Request) {
		// Logic to verify entered username and password from users.json?
		fmt.Println("here")
		fmt.Println(r.ParseForm())
	})	


	http.HandleFunc("/signup" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/signup.html"))

		if err := templates.ExecuteTemplate(w, "signup.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	


	http.HandleFunc("/feed" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/feed.html"))

		if err := templates.ExecuteTemplate(w, "feed.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	

	http.HandleFunc("/api/feed" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("data/posts.json"))
 
		if err := templates.ExecuteTemplate(w, "posts.json", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/users" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/users.html"))

		if err := templates.ExecuteTemplate(w, "users.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	

	http.HandleFunc("/api/users" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("data/users.json"))
 
		if err := templates.ExecuteTemplate(w, "users.json", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil));
}