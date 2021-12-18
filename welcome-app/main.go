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


//Go application entrypoint
func main() {

	curruser := CurrUser{"vihaha"}
	// myposts := Post{5, "test", "test", "test", "test"}
	currposts := []feed.Post{}

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
			templates := template.Must(template.ParseFiles("templates/welcome.html"))

			if err := templates.ExecuteTemplate(w, "welcome.html", curruser); err != nil {
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
		currposts = []feed.Post{}
		log.Println(curruser)
		log.Println(response)
		for i,_ := range response.Postid {
			currposts = append(currposts, feed.Post{
				PostID: int(response.Postid[i]),
				Title: response.Title[i],
				Author: response.Author[i],
				Description: response.Description[i],
				Timestamp: response.Timestamp[i],
			})
		}



		// for _, p := range response.FeedData {
		// 	posts = append(posts, feed.Posts{
		// 		PostID:      int(p.Postid),
		// 		Title:       p.Title,
		// 		Author:      p.Author,
		// 		Description: p.Description,
		// 		Timestamp:   p.Timestamp})
		// }
		// log.Println(response.postid)


		templates := template.Must(template.ParseFiles("templates/feed.html"))

		if err := templates.ExecuteTemplate(w, "feed.html", currposts); err != nil {
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

	http.HandleFunc("/signout", func(w http.ResponseWriter, r *http.Request) {

		// Reset user to guest
		curruser.Username = "Guest"

		// Redirect to homepage
		templates := template.Must(template.ParseFiles("templates/welcome.html"))

		if err := templates.ExecuteTemplate(w, "welcome.html", curruser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
