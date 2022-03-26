// need to write inside struct is call method

package main

import "fmt"

type User struct {
	Name, Email string
	Age         int
	Status      bool
}

func (u User) GetStatus() {
	fmt.Println("is user Active: ", u.Status)
}
func (u *User) ChangeEmail() {
	u.Email = "change@gemail.com"
}
func main() {
	mahbub := User{"Mahbub", "m@gmail.com", 25, true}
	tonmoy := User{"Mahbub", "m@gmail.com", 25, false}
	mahbub.GetStatus()
	tonmoy.GetStatus()
	fmt.Println(mahbub)
	mahbub.ChangeEmail()
	fmt.Println(mahbub)
}
