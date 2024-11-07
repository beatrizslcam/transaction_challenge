package user

import (
	"fmt"
	"transactions/repository"

	"golang.org/x/crypto/bcrypt"
)
	
	type ManageLogin struct{
		Repo repository.Repository
	}

type Users interface{
	IsAuthenticated(string)bool
	CreatePassword(string) ([]byte, error)
}

type User struct{
	ID string

}


func CreatePassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err!=nil {
		return nil, fmt.Errorf("error when hashing password due to: %v", err)
	}

	return hashedPassword, nil
}

