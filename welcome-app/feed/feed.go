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

type Post struct {
	PostID      int
	Title       string
	Author      string
	Description string
	Timestamp   string
}

func (s *Server) GetFeed(ctx context.Context, in *FeedRequest) (*FeedResponse, error) {
	log.Printf("Receieved following details from Client: \nusername: %s", in.Username)
	// Open our jsonFile
	jsonFile, err := os.Open("../data/users.json") //modified temporarily for testing

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened users.json")
	}
	

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

	jsonFile, err = os.Open("../data/posts.json") //modified temporarily for testing

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened posts.json")
	}
	

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ = ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var posts []Post
	filterPostIDs := []int32{}
	filterPostTitles := []string{}
	filterPostAuthor := []string{}
	filterPostDescription := []string{}
	filterPostTimestamps := []string{}
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &posts)
	fmt.Println(posts)

	for _, p := range posts {
		// fmt.Println(b.Username)
		// fmt.Println(b.Password)
		if followers[p.Author] {
			filterPostIDs = append(filterPostIDs, int32(p.PostID))
			filterPostTitles = append(filterPostTitles, p.Title)
			filterPostAuthor = append(filterPostAuthor, p.Author)
			filterPostDescription = append(filterPostDescription, p.Description)
			filterPostTimestamps = append(filterPostTimestamps, p.Timestamp)
			fmt.Println("Filtered Posts for feed Successfully!")
		}
	}
	fmt.Println(filterPostIDs, filterPostTitles)

	return &FeedResponse{Postid: filterPostIDs,
		Title:       filterPostTitles,
		Author:      filterPostAuthor,
		Description: filterPostDescription,
		Timestamp:   filterPostTimestamps}, nil
}
func (s *Server) PostToServer(ctx context.Context, in *PostData) (*PostDataResponse, error) {
	jsonFile, err := os.Open("../data/posts.json") //modified temporarily for testing

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened posts.json")
	}
	

	// defer the closing of our jsonFile so that we can parse it later on

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var posts []Post
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &posts)
	jsonFile.Close()
	posts = append(posts, Post{
		PostID:      int(in.Postid),
		Title:       in.Title,
		Author:      in.Author,
		Description: in.Description,
		Timestamp:   in.Timestamp,
	})
	byteValue, err = json.Marshal(posts)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("../data/posts.json", byteValue, 0644) //modified temporarily for testing
	if err != nil {
		log.Fatalf("failed to write to file on server: %v", err)
	}
	return &PostDataResponse{Success: true}, nil
}
