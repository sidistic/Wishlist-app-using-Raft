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
	//Instantiate a Welcome struct object and pass in some random information.
	//We shall get the name of the user as a query parameter from the URL
	welcome := Welcome{"Viha", time.Now().Format(time.Stamp)}

	//We tell Go exactly where we can find our html file. We ask Go to parse the html file (Notice
	// the relative path). We wrap it in a call to template.Must() which handles any errors 
	// and halts if there are fatal errors



	//Our HTML comes with CSS that go needs to provide when we run the app. Here we tell go to create
	// a handle that looks in the static directory, go then uses the "/static/" as a url that our
	//html can refer to when looking for our css and other files.

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")))) //Go looks in the relative static directory first, then matches it to a
	// 		//url of our choice as shown in http.Handle("/static/"). 
    //   //This url is what we need when referencing our css files
	// 		//once the server begins. Our html code would therefore be <link rel="stylesheet"  href="/static/stylesheet/...">
	// 		//It is important to note the final url can be whatever we like, so long as we are consistent.

	//This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	http.HandleFunc("/welcome" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/welcome.html"))

		//Takes the name from the URL query e.g ?name=Martin, will set welcome.Name = Martin.
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name;
		}

		//If errors show an internal server error message
		//I also pass the welcome struct to the welcome-template.html file. 
		if err := templates.ExecuteTemplate(w, "welcome.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Login Page
	http.HandleFunc("/login" , func(w http.ResponseWriter, r *http.Request) {

		// Parse login.html
		templates := template.Must(template.ParseFiles("templates/login.html"))

		//If errors show an internal server error message
		//I also pass the welcome struct to the welcome-template.html file. 
		if err := templates.ExecuteTemplate(w, "login.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	

	// Login Page
	http.HandleFunc("/api/feed" , func(w http.ResponseWriter, r *http.Request) {

		// Parse login.html
		templates := template.Must(template.ParseFiles("data/posts.json"))

		//If errors show an internal server error message
		//I also pass the welcome struct to the welcome-template.html file. 
		if err := templates.ExecuteTemplate(w, "posts.json", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})		

	http.HandleFunc("/verify-login" , func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("here")
		fmt.Println(r.ParseForm())
		// Logic to verify entered username and password from users.json?
	})	

	// Signup Page
	http.HandleFunc("/signup" , func(w http.ResponseWriter, r *http.Request) {

		// Parse login.html
		templates := template.Must(template.ParseFiles("templates/signup.html"))

		//If errors show an internal server error message
		//I also pass the welcome struct to the welcome-template.html file. 
		if err := templates.ExecuteTemplate(w, "signup.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	

	// Signup Page
	http.HandleFunc("/feed" , func(w http.ResponseWriter, r *http.Request) {

		// Parse login.html
		templates := template.Must(template.ParseFiles("templates/feed.html"))

		//If errors show an internal server error message
		//I also pass the welcome struct to the welcome-template.html file. 
		if err := templates.ExecuteTemplate(w, "feed.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt
	fmt.Println(http.ListenAndServe(":8080", nil));
}