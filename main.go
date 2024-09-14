package main

import (
	//"crypto/internal/boring/sig"
	"net/http"
	//"net/url"
	"html/template"
	"github.com/progressive-newbie263/signup-login-session/users"
)

func getLogInPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "log-in.html", nil)
}

func getSignUpPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "sign-up.html", nil)
}

func templating(w http.ResponseWriter, fileName string, data interface{}) {
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, data)
}

func logInUser(w http.ResponseWriter, r *http.Request) {

}

func signUpUser(w http.ResponseWriter, r *http.Request) {

}

// type User struct {
// 	Email string
// 	Password string
// }

func getUser(r *http.Request) users.User{
	email := r.FormValue("email")
	password := r.FormValue("Password")
	return users.User {
		Email: email,
		Password: password,
	}
}

func userHandler (w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	//login session
	case "log-in":
		logInUser(w, r)
	case "sign-up": 
		signUpUser(w, r)
	case "/log-in-form":
		getLogInPage(w, r)

	//sign up session.
	case "/sign-up-form":
		getSignUpPage(w, r)
	}
}

func main() {
	http.HandleFunc("/", userHandler)
	http.ListenAndServe("", nil)
}