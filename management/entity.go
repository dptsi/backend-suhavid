package management

import "fmt"

type User struct {
	ID        int
	FisrtName string
	LastName  string
	Email     string
	IsActive  bool
}

func (user User) displayUser() string {
	result := fmt.Sprintf("Name : %s %s, Email : %s, user.FirstName, user.LastName, user.Email")
	return result
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

// Embeded struct
func (group Group) DisplayGroup() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))

	fmt.Println("User name :")
	for _, user := range group.Users {
		fmt.Println(user.FisrtName)
	}
}
