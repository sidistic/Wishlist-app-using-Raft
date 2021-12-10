package main

import (
	"net/http"
	"fmt"
	"html/template"
)

//Create a struct that holds information to be displayed in our HTML file
type User struct {
	Name string
}

//Go application entrypoint
func main() {

	user := User{"Guest"}


	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))


	http.HandleFunc("/welcome" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/welcome.html"))

		if err := templates.ExecuteTemplate(w, "welcome.html", user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})


	http.HandleFunc("/login" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/login.html"))

		if err := templates.ExecuteTemplate(w, "login.html", user); err != nil {
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

		if err := templates.ExecuteTemplate(w, "signup.html", user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	


	http.HandleFunc("/feed" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/feed.html"))

		if err := templates.ExecuteTemplate(w, "feed.html", user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	

	http.HandleFunc("/api/feed" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("data/posts.json"))
 
		if err := templates.ExecuteTemplate(w, "posts.json", user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/users" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/users.html"))

		if err := templates.ExecuteTemplate(w, "users.html", user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	

	http.HandleFunc("/api/users" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("data/users.json"))
 
		if err := templates.ExecuteTemplate(w, "users.json", user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil));
}