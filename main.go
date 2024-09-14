package main

import (
	"net/http"
	"html/template"
	"github.com/progressive-newbie263/signup-login-session/users"
)

func getLogInPage(w http.ResponseWriter, _ *http.Request) {
	templating(w, "log-in.html", nil)
}

func getSignUpPage(w http.ResponseWriter, _ *http.Request) {
	templating(w, "sign-up.html", nil)
}

func templating(w http.ResponseWriter, fileName string, data interface{}) {
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, data)
}

//login function
func logInUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	ok := users.DefaultUserService.VerifyUser(newUser)
	
	// if signup failed
	if !ok {
		fileName := "log-in.html"
		t, _ := template.ParseFiles(fileName)
		t.ExecuteTemplate(w, fileName, "User log in failed.")
		return
	}

	//if sign up success
	fileName := "log-in.html"
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, "user logged in successfully.")
	return
}

//signup function
func signUpUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	err := users.DefaultUserService.CreateUser(newUser)
	
	// if signup failed
	if err != nil {
		fileName := "sign-up.html"
		t, _ := template.ParseFiles(fileName)
		t.ExecuteTemplate(w, fileName, "New user sign-up failed.")
		return
	}

	//if sign up success
	fileName := "sign-up.html"
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, "New user sign-up successfully.")
	return
}	

func getUser(r *http.Request) users.User{
	//this one is really case sensitive!
	//email and password has to be lowercase as html forms.
	//you can use it depends on the form though!
	email := r.FormValue("email")
	password := r.FormValue("password") 

	return users.User {
		Email: email,
		Password: password,
	}
}

func userHandler (w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	//login session
	case "/log-in":
		logInUser(w, r)
	case "/sign-up": 
		signUpUser(w, r)
	//log in session html
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