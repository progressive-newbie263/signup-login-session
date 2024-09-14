package users

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

/*
	one thing to remember. When storing user's username and password
	we don't want to save them directly into the database, but we want to
	save them as a hash in the database instead.

	Reason is because when we create account, if we use a password for an email,
	we likely will be using that same password for other email as well.

	So, if the hacker was able to attack the database, they will be able to use the
	combination of password and account infos to break others social platform accounts
	of that email as well.
*/
type User struct {
	Email string
	Password string
}

//so for that reason, we will create a seperate database:
type authUser struct {
	email string
	passwordHash string
}

var authUserDB = map[string]authUser{} //email => authUser{email, hash}

var DefaultUserService userService

type userService struct {

}

func (userService) VerifyUser(user User) bool {
	authUser, ok := authUserDB[user.Email]
	if !ok {
		return false
	}
	err := bcrypt.CompareHashAndPassword(
		[]byte(authUser.passwordHash),
		[]byte(user.Password))
	
	return err == nil
}


func (userService) CreateUser(newUser User) error{
	//if the user already exists
	_, ok := authUserDB[newUser.Email]
	//this has to be "if ok". Since you can only show the user not exists
	// after creating an user
	if ok {
		fmt.Println("user already exists")
		return errors.New("user already exists")
	}

	passwordHash, err := getPasswordHash(newUser.Password)

	if err != nil {
		fmt.Println("Get password hash")
		return err
	}

	newAuthUser := authUser {
		email: newUser.Email,
		passwordHash: passwordHash,
	}

	authUserDB[newAuthUser.email] = newAuthUser
	return nil
}

//so, to create password hash, we are using bcrypt library:
func getPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(hash), err
}