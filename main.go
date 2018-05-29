package main

import (
	"html/template"
	"net/http"
	"fmt"
	"time"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/gorilla/mux"
	"github.com/gorilla/context"
	"math/rand"
)

type datastore struct {
	MySql *sqlx.DB
}


var ds datastore

func InitMySqlConn(dsn string) (*sqlx.DB, error) {
	var err error
	ds.MySql, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	ds.MySql.SetMaxOpenConns(50)

	ds.MySql.SetMaxIdleConns(20)
	ds.MySql.SetConnMaxLifetime(time.Hour)

	if err = ds.MySql.Ping(); err != nil {
		return nil, err
	}

	return ds.MySql, nil
}

func main() {
	//mysqlDsn :=  os.Getenv("MYSQL_DSN")
	mysqlDsn := "%s:%s@tcp(%s:%s)/%s?timeout=30s&readTimeout=1s&writeTimeout=1s"
	mysqlDsn = fmt.Sprintf(
		mysqlDsn,
		os.Getenv("MYSQL_ROOT"),
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))

	fmt.Printf("Initializing MySql connection to: %s\n", mysqlDsn)

	trials, maxTrials := 0, 15 // TODO import maxTrials value from a config file

	for {
		db, err := InitMySqlConn(mysqlDsn)
		if err != nil {
			log.Printf("Unable to connect to MySql (trial %d): %s\n", trials, err)
			time.Sleep(time.Duration(1) * time.Second)
			trials++
			if trials >= maxTrials {
				os.Exit(1)
			}
		} else {
			fmt.Printf("Connected to db\n")
			defer db.Close()
			break
		}
	}

	adminserver := mux.NewRouter()

	adminserver.HandleFunc("/", LoginHandler)


	adminserver.Handle("/add", AuthMiddleware(http.HandlerFunc(AddPartnerHandler)))
	adminserver.Handle("/partners", AuthMiddleware(http.HandlerFunc(ListPartnersHandler)))
	adminserver.Handle("/delete/{id:[0-9]+}", AuthMiddleware(http.HandlerFunc(DeletePartnerHandler)))
	adminserver.Handle("/edit/{id:[0-9]+}", AuthMiddleware(http.HandlerFunc(EditPartnerHandler)))

	adminserver.Handle("/users", AuthMiddleware(http.HandlerFunc(ListUsersHandler)))

	adminserver.HandleFunc("/register", RegisterHandler)
	adminserver.HandleFunc("/login", LoginHandler)
	adminserver.HandleFunc("/logout", LogoutHandler)

	adminserver.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	//adminserver.PathPrefix("/static/").Handler(http.StripPrefix("/static/", imgServer))

	http.Handle("/", adminserver)

	srvPort := "3001" // TODO declare an ADSERVER_PORT env variable and use it here
	fmt.Printf("Server starting on port :%s\n", srvPort)
	log.Fatal(http.ListenAndServe(":"+srvPort, adminserver))
}

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("e")
		if err != nil {
			fmt.Printf("Error at cookie:  %v \n", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		decryptedEmail, err := decrypt(c.Value)
		if err != nil {
			fmt.Printf("Error at decrypting:  %v \n", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		selectAUserQuery := "SELECT * FROM Users WHERE email=?"
		var u User
		err = ds.MySql.Get(&u, selectAUserQuery, decryptedEmail)
		if err != nil {
			fmt.Printf("Error not found user:  %v \n", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(100-0) + 0

		context.Set(r, "user", u)
		context.Set(r, "random", random)

		h.ServeHTTP(w, r)
	})
}

func execTmpl(w http.ResponseWriter, data interface{}, file string){
	tmpl := template.Must(template.New("").ParseFiles("./tpl/navbar.html",file, "./tpl/base.html", "./tpl/sidebar.html"))
	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		fmt.Printf("Cannot execute template %s \n", err)
		return
	}
}

