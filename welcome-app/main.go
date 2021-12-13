package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"welcome-app/feed"
	"welcome-app/login"

	"google.golang.org/grpc"
)

//Create a struct that holds information to be displayed in our HTML file
type CurrUser struct {
	Username string
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
	posts := []feed.Posts{}

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/welcome.html"))

		if err := templates.ExecuteTemplate(w, "welcome.html", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/login.html"))

		if err := templates.ExecuteTemplate(w, "login.html", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/verify-login", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		var ip_username = r.Form["username"][0]
		var ip_password = r.Form["password"][0]

		var conn *grpc.ClientConn
		conn, err := grpc.Dial(":9000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		l := login.NewAuthServiceClient(conn)

		response, err := l.Authenticate(context.Background(), &login.LoginDetails{Username: ip_username, Password: ip_password})
		if err != nil {
			log.Fatalf("Error when calling Authenticate: %s", err)
		}
		log.Printf("Response from server: Name: %s", response.Name)

		if response.Done {
			curruser.Username = response.Name
			templates := template.Must(template.ParseFiles("templates/feed.html"))

			if err := templates.ExecuteTemplate(w, "feed.html", curruser); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			templates := template.Must(template.ParseFiles("templates/login.html"))

			if err := templates.ExecuteTemplate(w, "login.html", curruser); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/signup.html"))

		if err := templates.ExecuteTemplate(w, "signup.html", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/feed", func(w http.ResponseWriter, r *http.Request) {
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(":9000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		f := feed.NewFeedServiceClient(conn)

		response, err := f.GetFeed(context.Background(), &feed.FeedRequest{Username: curruser.Username})
		if err != nil {
			log.Fatalf("Error when calling GetFeed: %s", err)
		}
		posts = []feed.Posts{}
		log.Println(response.FeedData)
		for _, p := range response.FeedData {
			posts = append(posts, feed.Posts{
				PostID:      int(p.Postid),
				Title:       p.Title,
				Author:      p.Author,
				Description: p.Description,
				Timestamp:   p.Timestamp})
		}

		templates := template.Must(template.ParseFiles("templates/feed.html"))

		if err := templates.ExecuteTemplate(w, "feed.html", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	http.HandleFunc("/api/feed", func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("data/posts.json"))

		if err := templates.ExecuteTemplate(w, "posts.json", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/add-post", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("templates/users.html"))

		if err := templates.ExecuteTemplate(w, "users.html", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("data/users.json"))

		if err := templates.ExecuteTemplate(w, "users.json", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
