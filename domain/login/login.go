package login

import (
	"fmt"
	"transactions/repository"
	"transactions/service"

	"golang.org/x/crypto/bcrypt"
)
	
	type ManageLogin struct{
		Repo repository.RepositoryAccount
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




func (ml *ManageLogin)Login(cpf string, password string)(string,error){
	
	found, userAccount := ml.Repo.FindAccountByCpf(cpf)
	if !found {
		return "", fmt.Errorf("account doesn't exist")
	} 

	//Validate Password
	err := bcrypt.CompareHashAndPassword([]byte(userAccount.Secret),[]byte(password))
	if err!=nil {
		return "", fmt.Errorf("error comparing password: %v", err)
	}

	token, err:= ml.Auth.GenerateToken(userAccount.ID)
	if err!= nil{
		return "", fmt.Errorf("error generating token")
	}

	return token, nil
	


}

