package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	// "os/exec"

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

	resp, err := http.Get("http://127.0.0.1:12380/users")
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// we initialize our Users array
	var users []Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(body, &users)

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
