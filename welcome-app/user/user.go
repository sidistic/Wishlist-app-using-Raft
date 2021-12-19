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

	file, err := ioutil.ReadFile("data/users.json")
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
	err = ioutil.WriteFile("data/users.json", dataBytes, 0644)
	if err != nil {
		log.Fatalf("failed to write to file on server: %v", err)
	}
	return &SignUpResponse{Success: true}, nil
}
