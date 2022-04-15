package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)

func main() {
	password, err := hashPassword("VINIcius@golang")	

	fmt.Println(password)
	fmt.Println(err)
}

func hashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(bytes), []byte("VINIcius@golang")))

	return string(bytes), err
}