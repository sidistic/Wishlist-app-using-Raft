package feed

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/net/context"
)

type Server struct {
}

type Users struct {
	Username string
	Name     string
	Password string
	Follows  []string
}

type Posts struct {
	PostID      int
	Title       string
	Author      string
	Description string
	Timestamp   string
}

func (s *Server) GetFeed(ctx context.Context, in *FeedRequest) (*FeedResponse, error) {
	log.Printf("Receieved following details from Client: \nusername: %s", in.Username)
	// Open our jsonFile
	jsonFile, err := os.Open("../data/users.json")

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

	followers := map[string]bool{}

	for _, b := range users {
		// fmt.Println(b.Username)
		// fmt.Println(b.Password)
		if b.Username == in.Username {
			for _, follower := range b.Follows {
				followers[follower] = true
			}
			fmt.Println("Found followers Successfully!")
		}
	}

	fmt.Println(in.Username)
	fmt.Println(followers)

	// curruser := CurrUser{ip_username}

	// fmt.Println(curruser)

	jsonFile, err = os.Open("../data/posts.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened posts.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ = ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var posts []Posts
	var filterPosts []*Post

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &posts)

	for _, p := range posts {
		// fmt.Println(b.Username)
		// fmt.Println(b.Password)
		if followers[p.Author] {
			filterPosts = append(filterPosts, &Post{
				Postid:      int32(p.PostID),
				Title:       p.Title,
				Author:      p.Author,
				Description: p.Description,
				Timestamp:   p.Timestamp})
			fmt.Println("Filtered Posts for feed Successfully!")
		}
	}

	return &FeedResponse{FeedData: filterPosts}, nil
}
