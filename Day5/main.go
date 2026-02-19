package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName string
	Password string // it should be stored as hashed password
}

type UserMap struct {
	Data map[string]*User
}

func NewUser() *UserMap {
	return &UserMap{
		Data: make(map[string]*User), // initialize map
	}
}

// making method to append users to the system

func (mp *UserMap) addUser(userData User) {

	// convert password into hash

	hashPassword, err := HashPassword(userData.Password)
	// In this function we simply add that mp element to the Data
	if err != nil {
		fmt.Println("Error storing user data")
		return

	}

	mp.Data[userData.UserName] = &User{
		UserName: userData.UserName,
		Password: hashPassword,
	} // storing address to user

	fmt.Printf("User %v added successfully\n", userData.UserName)

}

// in this method we simply check whether the user is authenticated or not

func (mp *UserMap) checkAuthentication(userData User) {
	val, ok := mp.Data[userData.UserName] // checking our username is authenticated or not

	if !ok { // means it is not authenticated
		fmt.Printf("User %v not found\n", userData.UserName)
		return
	}

	// now compare password if password matches with hash then ok

	userHash := val.Password // this contain hash password for this user

	flag := CheckPasswordHash(userData.Password, userHash)

	if !flag {
		fmt.Printf("User %v password is not correct\n", userData.UserName)
		return
	}

	fmt.Printf("User %v password is correct . Successfully Authenticated\n", userData.UserName)
}

// this function hashes my password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// making user login system
func main() {

	fmt.Println("Making User login system ")

	mp := NewUser()

	//fmt.Println(mp)
	//fmt.Println(mp.Data)
	Users := []User{
		{UserName: "alice", Password: "pass123"},
		{UserName: "bob", Password: "bob@321"},
		{UserName: "charlie", Password: "charlie12"},
		{UserName: "david", Password: "david!45"},
		{UserName: "emma", Password: "emma2024"},
	}

	mp.addUser(Users[0])
	mp.addUser(Users[1])
	mp.addUser(Users[2])

	mp.checkAuthentication(Users[0]) // for checking authentication
	mp.checkAuthentication(Users[1])
	mp.checkAuthentication(User{"alice", "pass122"}) // wrong password
	mp.checkAuthentication(Users[3])                 // no entry

}
