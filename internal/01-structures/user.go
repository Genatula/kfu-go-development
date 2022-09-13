package structures

import (
	"fmt"
	"strings"
)

type User struct {
	Firstname   string
	Lastname    string
	Patronymic  string
	PhoneNumber string
	Age         int8
}

func (user *User) ChangePhoneNumber(newNumber string) {
	if strings.Contains(newNumber, "+7") {
		user.PhoneNumber = newNumber
	} else if newNumber[0:1] == "7" || newNumber[0:1] == "8" {
		var length = len(newNumber)
		var number = "+7" + newNumber[1:length]
		user.PhoneNumber = number
	} else {
		var number = "+7" + newNumber
		user.PhoneNumber = number
	}
}

func (user *User) PrintUserInfo() {
	fmt.Printf("User = {firstname: %s, lastname: %s, patronymic: %s, PhoneNumber: %s, age: %d}", user.Firstname,
		user.Lastname, user.Patronymic, user.PhoneNumber, user.Age)
}
