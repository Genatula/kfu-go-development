package main

import "kfu-go-development/internal/01-structures"

func main() {
	var user = structures.User{"Ilnas", "Genatulin", "Nailevich", "+71234567890", 20}
	user.PrintUserInfo()
	user.ChangePhoneNumber("2345678901")
	user.PrintUserInfo()
}
