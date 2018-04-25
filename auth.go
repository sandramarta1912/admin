package main

import (
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"github.com/sudo-suhas/symcrypto"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := partnerTemplates.ExecuteTemplate(w, "register", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}

		name := r.PostFormValue("name")
		email := r.PostFormValue("email")

		password := r.PostFormValue("password")
		bpassword := []byte(password)
		hashedPassword, err := bcrypt.GenerateFromPassword(bpassword, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(hashedPassword))

		// Comparing the password with the hash
		err = bcrypt.CompareHashAndPassword(hashedPassword, bpassword)
		if err != nil {
			panic(err)
		}

		var createUserQuery = "INSERT Users SET name=?,email=?, password=?"
		_, err = ds.MySql.Exec(createUserQuery, name, email, hashedPassword)
		if err != nil {
			panic(err)
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)

		}
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c, err := r.Cookie("e")
		if err != nil {
			fmt.Printf("Error at cookie:  %v \n", err)
			execLoginTmpl(w, r)
			return
		}
		decryptedEmail, err := decrypt(c.Value)
		if err != nil {
			fmt.Printf("Error at decrypting:  %v \n", err)
			execLoginTmpl(w, r)
			return
		}
		selectAUserQuery := "SELECT * FROM Users WHERE email=?"
		var u User
		err = ds.MySql.Get(&u, selectAUserQuery, decryptedEmail)
		if err != nil {
			fmt.Printf("Error not found user:  %v \n", err)
			execLoginTmpl(w, r)
			return
		}

		http.Redirect(w, r, "/partners", http.StatusSeeOther)
	}

	if r.Method == http.MethodPost {
		fmt.Println("In login post")

		err := r.ParseForm()
		if err != nil {
			panic(err)
		}

		email := r.PostFormValue("email")


		password := r.PostFormValue("password")

		bpassword := []byte(password)

		selectAUserQuery := "SELECT * FROM Users WHERE email=?"
		var u User
		err = ds.MySql.Get(&u, selectAUserQuery, email)
		if err != nil {
			fmt.Printf("Cannot get user form db %s \n",err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(u.Password), bpassword)
		if err != nil {
			fmt.Printf("Bad password %s \n", err)
			return
		} else{
			encryptedEmail, _ := encrypt(email)

			cookie := &http.Cookie{Name:"e", Value:encryptedEmail, Path:"/"}

			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/partners", http.StatusSeeOther)
		}
	}
}
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "e",
		MaxAge: -1}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}

var crypto, _ = symcrypto.New("sdfdbhjerrzxcsidufsdfsdgsdwhgtyu")

func encrypt(str string) (string, error){
	encrypted, err := crypto.Encrypt(str)
	if err != nil {
		fmt.Printf("encrypt %s", err)
		return "", err

	}
	return encrypted, nil
}

func decrypt(str string) (string, error){
	decrypted, err := crypto.Decrypt(str)
	if err != nil {
		fmt.Printf("decrypt %s", err)
		return "", err

	}
	return decrypted, nil
}

func execLoginTmpl(w http.ResponseWriter, r *http.Request){
	err := partnerTemplates.ExecuteTemplate(w, "login", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}