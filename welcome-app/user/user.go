package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
)

type User struct {
	Username string
	Name     string
	Password string
	Follows  []string
}

type Server struct {
}

func (s *Server) SignUpUser(ctx context.Context, in *SignUpRequest) (*SignUpResponse, error) {
	if in.Password != in.ConfirmPassword {
		return &SignUpResponse{Success: false}, nil
	}

	file, err := ioutil.ReadFile("data/users.json") //modified temporarily for testing
	if err != nil {
		fmt.Println(err)
	}

	users := []User{}

	json.Unmarshal(file, &users)

	for _, u := range users {
		if u.Username == in.Username {
			return &SignUpResponse{Success: false}, nil
		}
	}

	users = append(users, User{
		Username: in.Username,
		Name:     in.Name,
		Password: in.Password,
		Follows:  []string{},
	})

	dataBytes, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("data/users.json", dataBytes, 0644) //modified temporarily for testing
	if err != nil {
		log.Fatalf("SignUpUser: failed to write to file on server: %v", err)
	}
	return &SignUpResponse{Success: true}, nil
}

func (s *Server) GetFollowing(ctx context.Context, in *FollowerRequest) (*FollowerResponse, error) {
	file, err := ioutil.ReadFile("data/users.json") //modified temporarily for testing
	if err != nil {
		fmt.Println(err)
	}

	users := []User{}
	notFollowed := []string{}
	Followed := []string{}

	json.Unmarshal(file, &users)

	for _, u := range users {
		if in.Username == u.Username {
			Followed = u.Follows
		}
	}
	for _, u := range users {
		if !CheckIfFollowed(u.Username, Followed) {
			notFollowed = append(notFollowed, u.Username)
		}
	}

	return &FollowerResponse{
		Followers:    Followed,
		NotFollowers: notFollowed,
	}, nil
}

func CheckIfFollowed(username string, follows []string) bool {
	for _, s := range follows {
		if s == username {
			return true
		}
	}
	return false
}

func (s *Server) UpdateFollower(ctx context.Context, in *UpdateFollowersRequest) (*UpdateFollowersResponse, error) {

	file, err := ioutil.ReadFile("data/users.json") //modified temporarily for testing
	if err != nil {
		fmt.Println(err)
	}

	users := []User{}

	json.Unmarshal(file, &users)

	for i, u := range users {
		if u.Username == in.Username {
			if in.IsFollow {
				users[i].Follows = append(users[i].Follows, in.Newuser)
			} else {
				NewFollows := []string{}
				for _, f := range u.Follows {
					if in.Newuser != f {
						NewFollows = append(NewFollows, f)
					}
				}
				users[i].Follows = NewFollows
			}
		}
	}

	dataBytes, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("data/users.json", dataBytes, 0644) //modified temporarily for testing
	if err != nil {
		log.Fatalf("UpdateFollower: failed to write to file on server: %v", err)
	}
	return &UpdateFollowersResponse{Success: true}, nil
}
