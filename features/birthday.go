package features

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Users is a struct that contains a slice of the User struct. It is made to hold the data from all of the users in the JSON file
type Users struct {
	Users []User `json:"users"`
}

// User is a struct that contains an Id field and a Birthdate field. This is the container that holds the data of one user from the JSON file
type User struct {
	Id        string `json:"id"`
	Birthdate string `json:"birthday"`
}

// ReadBirthdayData initializes a Users struct and reads the data from the birthdayData JSON file. It then returns the Users struct and an error
func ReadBirthdayData() (Users, error) {
	var users Users

	jsonFile, err := os.Open("data/birthdayData.json")
	if err != nil {
		fmt.Println(err)
	} else {
    fmt.Println("Successfully opened birthdayData")
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)  

  return users, json.Unmarshal(byteValue, &users)
}

func PrintBirthdayData(users Users) {
	for i := 0; i < len(users.Users); i++ {
    fmt.Println("User id: " + users.Users[i].Id)
    fmt.Println("Birthdate: " + users.Users[i].Birthdate)
	}
}

func AddBirthday() {



}
