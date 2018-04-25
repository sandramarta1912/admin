package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/context"
)

type userCollection struct {
	Users []User
}

type User struct {
	Id int `db:"id"`
	Name string `db:"name"`
	Email string `db:"email"`
	Password string `db:"password"`
}
type usersData struct {
	Random int
	User User
	Users userCollection
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	userContext := context.Get(r, "user")
	randomContext := context.Get(r, "random")

	if r.Method == http.MethodGet {
		var users userCollection
		var selectAllUsersQuery = "SELECT * FROM Users"
		err := ds.MySql.Select(&users.Users, selectAllUsersQuery)
		if err != nil {
			fmt.Printf("Cannot get user from db %s \n", err)
			return
		}

		var random int
		var user User
		switch r := randomContext.(type) {
		case int:
			random = r

		}
		switch u := userContext.(type) {
		case User:
			user = u

		}
		execTmpl(w, usersData{random,  user, users}, "./tpl/all_users.html")

	}
}
