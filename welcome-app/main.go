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
	Username string
}

type Users struct {
    Username   string 
    Name   string 
    Password    string    
    Follows  []string
}

// type User struct {
//     Username   string 
//     Name   string 
//     Password    string    
//     Follows  []string
// }

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

		r.ParseForm()

		var ip_username = r.Form["username"][0]
		var ip_password = r.Form["password"][0]

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
		var users []Users

		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		json.Unmarshal(byteValue, &users)

		for _, b := range users {
			// fmt.Println(b.Username)
			// fmt.Println(b.Password)
			if b.Username == ip_username && b.Password == ip_password {
				// validate = true
				fmt.Println("Validated Successfully!")
			}
		}

		fmt.Println(ip_username)
		fmt.Println(ip_password)

		curruser := CurrUser{ip_username}

		fmt.Println(curruser)

		// need to re route to /feed

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