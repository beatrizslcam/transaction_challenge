package login

import (
	"fmt"
	"transactions/repository"
	"transactions/service"

	"golang.org/x/crypto/bcrypt"
)
	
	type ManageLogin struct{
		Repo repository.Repository
		Auth service.Auth
	}

type Login interface{
	Login(string,string)string
	IsAuthenticated(string)bool
}



func (ml *ManageLogin) IsAuthenticated(token string)bool{
	isauth := ml.Auth.ValidateToken(token)
	return isauth
}




func (ml *ManageLogin)Login(cpf string, password string)string{
	found, userAccount := ml.Repo.FindAccountByCpf(cpf)
	if !found {
		return "Accoun't doesn't exists"
	} 
	err := bcrypt.CompareHashAndPassword([]byte(userAccount.Secret),[]byte(password))
		if err!=nil {
			fmt.Printf("Error comparing password: %v\n", err)
			return ""
		}

		token,err := service.GenerateToken(userAccount.ID)
		if err!= nil{
			fmt.Println("Error generating token")
			return ""
		}

		return token
	


}

