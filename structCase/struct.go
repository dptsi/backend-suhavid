package structCase

import "fmt"

type User struct {
	ID        int
	FisrtName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func main() {
	user := User{
		ID:        1,
		FisrtName: "Suhavid Hendra",
		LastName:  "Kusuma",
		Email:     "suhavidhendra@gmail.com",
		IsActive:  true,
	}

	user2 := User{2, "Agus Wahyudi", "Saputra", "agus@gmail.com", true}

	displayUser1 := displayUser(user)
	displayUser2 := displayUser(user2)

	fmt.Println(displayUser1)
	fmt.Println(displayUser2)

	//embeded struct
	users := []User{user, user2}

	group := Group{"Gamer", user, users, true}

	displayGroup(group)

}

// Struct sebagai parameter
func displayUser(user User) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s, user.FirstName, user.LastName, user.Email")
	return result
}

// Embeded struct
func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))

	fmt.Println("User name :")
	for _, user := range group.Users {
		fmt.Println(user.FisrtName)
	}
}
