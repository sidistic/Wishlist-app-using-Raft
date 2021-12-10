package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"html/template"
	"os"
	"io/ioutil"
)

//Create a struct that holds information to be displayed in our HTML file
type CurrUser struct {
	Name string
}

type Users struct {
    Users []User 
}

type User struct {
    Username   string 
    Name   string 
    Password    string    
    Follows  []string
}

//Go application entrypoint
func main() {

	curruser := CurrUser{"Guest"}


	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))


	http.HandleFunc("/welcome" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/welcome.html"))

		if err := templates.ExecuteTemplate(w, "welcome.html", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})


	http.HandleFunc("/login" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/login.html"))

		if err := templates.ExecuteTemplate(w, "login.html", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	


	http.HandleFunc("/verify-login" , func(w http.ResponseWriter, r *http.Request) {
		// Logic to verify entered username and password from users.json?
		fmt.Println("enter verify-login")

		r.ParseForm()
		for key, value := range r.Form {
			fmt.Printf("%s = %s\n", key, value)
		}

		// Open our jsonFile
		jsonFile, err := os.Open("data/users.json")

		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened users.json")

		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		// read our opened jsonFile as a byte array.
		byteValue, _ := ioutil.ReadAll(jsonFile)

		// we initialize our Users array
		var users Users

		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		json.Unmarshal(byteValue, &users)


		// we iterate through every user within our users array and
		// print out the user Type, their name, and their facebook url
		// as just an example

		fmt.Println(len(users.Users))
		// fmt.Println(len(Users))

		for i := 0; i < len(users.Users); i++ {
			fmt.Println("Username: " + users.Users[i].Username)
			fmt.Println("Name: " + users.Users[i].Name)
			fmt.Println("Password: " + users.Users[i].Password)
			// fmt.Println("Follows: " + users.Users[i].[].Follows)
		}		

		fmt.Println("end")

	})	


	http.HandleFunc("/signup" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/signup.html"))

		if err := templates.ExecuteTemplate(w, "signup.html", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	


	http.HandleFunc("/feed" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/feed.html"))

		if err := templates.ExecuteTemplate(w, "feed.html", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	

	http.HandleFunc("/api/feed" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("data/posts.json"))
 
		if err := templates.ExecuteTemplate(w, "posts.json", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/add-post" , func(w http.ResponseWriter, r *http.Request) {
		
	})	

	http.HandleFunc("/users" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/users.html"))

		if err := templates.ExecuteTemplate(w, "users.html", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	

	http.HandleFunc("/api/users" , func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("data/users.json"))
 
		if err := templates.ExecuteTemplate(w, "users.json", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil));
}