package birthday

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
  Displayname string `json:"displayname"`
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

// PrintBirthdayData will print all of the data from the Users struct to the console
func PrintBirthdayData(users Users) {
	for i := 0; i < len(users.Users); i++ {
    fmt.Println("User id: " + users.Users[i].Id)
    fmt.Println("Display name: " + users.Users[i].Id)
    fmt.Println("Birthdate: " + users.Users[i].Birthdate)
	}
}

// AddBirthday takes an integer id, and birthdate string and appends them to the json file. If the birthdate is already found, return an error
func AddBirthday(id string, displayname string, birthdate string) error {
  // 

  return nil // everything went right
}

// RemoveBirthday will remove the birthday of the user from the JSON file, if the user is not found return an error
func RemoveBirthday(id string) error {

  return nil // return nil if the function works as expected
}

// EditBirthday will remove an old birthday and add a new one, if an error is found with either operation it will return the error
func EditBirthday(id string, displayname string, birthdate string) error {
  var err error
  err = RemoveBirthday(id)
  if err != nil {
    // handle remove birthday error
  }

  err = AddBirthday(id, displayname, birthdate)
  if err != nil {
    // handle add birthday error
  }

  return nil
}

