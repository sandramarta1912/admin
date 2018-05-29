package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/context"
)

type managerCollection struct {
	Managers []Manager
}

type Manager struct {
	Id int `db:"id"`
	Name string `db:"name"`
	Email string `db:"email"`
	Password string `db:"password"`
}
type managersData struct {
	Random int
	Manager Manager
	Managers managerCollection
}

func ListManagersHandler(w http.ResponseWriter, r *http.Request) {
	managerContext := context.Get(r, "manager")
	randomContext := context.Get(r, "random")

	if r.Method == http.MethodGet {
		var managers managerCollection
		var selectAllManagersQuery = "SELECT * FROM Managers"
		err := ds.MySql.Select(&managers.Managers, selectAllManagersQuery)
		if err != nil {
			fmt.Printf("Cannot get manager from db %s \n", err)
			return
		}

		var random int
		var manager Manager
		switch r := randomContext.(type) {
		case int:
			random = r

		}
		switch m := managerContext.(type) {
		case Manager:
			manager = m

		}
		execTmpl(w, managersData{random,  manager, managers}, "./tpl/all_managers.html")

	}
}
