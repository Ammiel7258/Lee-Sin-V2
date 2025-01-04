package birthday

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Users is a struct that contains a slice of the User struct. It is made to hold the data from all of the users in the JSON file
type Users struct {
	Users []User `json:"users"`
}

// User is a struct that contains an Id field and a Birthdate field. This is the container that holds the data of one user from the JSON file
type User struct {
	Id          string `json:"id"`
	Displayname string `json:"displayname"`
	Birthdate   string `json:"birthday"`
}

var users Users

func init() {
	readBirthdayData()
}

// readBirthdayData reads the data from the birthdayData JSON file. It then returns the Users struct and an error
func readBirthdayData() error {
  jsonFile, err := os.Open("data/birthdayData.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	return json.NewDecoder(jsonFile).Decode(&users)
}

func writeBirthdayData() error {
  jsonFile, err := os.OpenFile("data/birthdayData.json", os.O_WRONLY|os.O_TRUNC, 0644)
  if err != nil {
    return err
  }
  defer jsonFile.Close()
  
  return json.NewEncoder(jsonFile).Encode(users)
}

// AddBirthday takes an integer id, and birthdate string and appends them to the json file. If the birthdate is already found, return an error
func AddBirthday(id string, displayname string, birthdate string) error {
	var user User
	user.Id = id
	user.Displayname = displayname
	user.Birthdate = birthdate

	for _, u := range users.Users {
		if u.Id == id {
      return errors.New("A user with this ID already exists!")
		}
	}

  users.Users = append(users.Users, user)

	return writeBirthdayData()
}

// RemoveBirthday will remove the birthday of the user from the JSON file, if the user is not found return an error
func RemoveBirthday(id string) error {
	var err error

	err = readBirthdayData()
	if err != nil {
		return err
	}

	return nil // return nil if the function works as expected
}

// EditBirthday will remove an old birthday and add a new one, if an error is found with either operation it will return the error
func EditBirthday(id string, displayname string, birthdate string) error {
	var err error

	err = RemoveBirthday(id)
	// handle remove birthday error
	if err != nil {
		return err
	}

	err = AddBirthday(id, displayname, birthdate)
	if err != nil {
		// handle add birthday error
		return err
	}

	return nil
}

// PrintBirthdayData will print all of the data from the Users struct to the console, returns nil if no errors were found
func PrintBirthdayData() error {
	err := readBirthdayData()
	if err != nil {
		return err
	}

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User id: " + users.Users[i].Id)
		fmt.Println("Display name: " + users.Users[i].Displayname)
		fmt.Println("Birthdate: " + users.Users[i].Birthdate + "\n")
	}
	return nil
}

func GetUsers() Users {
	return users
}
