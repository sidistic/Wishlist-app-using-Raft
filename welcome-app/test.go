package main

import "net/http"
import "fmt"
import "io/ioutil"
import "welcome-app/login"
import "encoding/json"


func main(){

	resp, err := http.Get("http://127.0.0.1:12380/users")
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println([]byte(string(body)))
	fmt.Println(string(body))
	var users []login.Users
	json.Unmarshal(body, &users)
	// if err := json.NewDecoder(resp.Body).Decode(users); err != nil {
	// 	// log.Fatalln(err)
	// }

	fmt.Println(users)
   
//Convert the body to type string
	// sb := string(body)
	// fmt.Printf(sb)
	// json.Unmarshal(body, &users)
	fmt.Println(users)
   
}