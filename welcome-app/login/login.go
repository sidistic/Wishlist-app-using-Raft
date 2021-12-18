package login

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

func (s *Server) Authenticate(ctx context.Context, in *LoginDetails) (*LoginResponse, error) {
	log.Printf("Receieved following details from Client: \nusername: %s\nPassword: %s ", in.Username, in.Password)
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

	done := false

	for _, b := range users {
		// fmt.Println(b.Username)
		// fmt.Println(b.Password)
		if b.Username == in.Username && b.Password == in.Password {
			// validate = true
			fmt.Println("Validated Successfully!")
			done = true
		}
	}

	fmt.Println(in.Username)
	fmt.Println(in.Password)

	// curruser := CurrUser{ip_username}

	// fmt.Println(curruser)
	return &LoginResponse{Name: in.Username, Done: done}, nil
}