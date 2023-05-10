package structCase

import (
	"backend-suhavid/management"
)

func main() {
	user := management.User{
		ID:        1,
		FisrtName: "Suhavid Hendra",
		LastName:  "Kusuma",
		Email:     "suhavidhendra@gmail.com",
		IsActive:  true,
	}

	user2 := management.User{2, "Agus Wahyudi", "Saputra", "agus@gmail.com", true}

	//embeded struct
	users := []management.User{user, user2}

	group := management.Group{"Gamer", user, users, true}

	group.DisplayGroup()

}
