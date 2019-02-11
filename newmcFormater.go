package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)

//Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name a type and a list of social links
type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Age  int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	jsonFile, err := os.Open("users.json")	//open up the jsonFile.
	// if os.Open returns an error then handle it.
	if err != nil {
		fmt.Println(err)
	}
	//read the opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)
	
	// Initialize Users array
	var users Users

	// Unmarshal the byteArray which contains the jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as jsut an example
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
	//	fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on.
	defer jsonFile.Close()
}
