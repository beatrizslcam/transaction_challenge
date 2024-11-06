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
	IsAuthenticated(token string)bool
}

type User struct{
	ID string

}



func (ml *ManageLogin) IsAuthenticated(token string)bool{
	return true
}
func CreatePassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err!=nil {
		return nil, fmt.Errorf("Error when hashing password due to: %v", err)
	}

	return hashedPassword, nil
}

