package account

import (
	"fmt"
	"transactions/repository"
) 



type ManageAccount struct{
	Repo repository.Repository
}


type  Account struct{
	ID string 
	Name  string 
	Cpf    string 
	Secret string 
	Balance  int 
	
}


func (ma *ManageAccount) CreateAccount(name string, cpf string) Account{
	does_exist,existing_account := ma.Repo.FindAccountByCpf(cpf)
		if does_exist {
			fmt.Printf("Account already exists")
			return existing_account.(Account)
		}
	
	new_account :=Account{
		Name:      name,
	Cpf:       cpf,
	Balance:   0,
	ID:        "ugiugiu",
	Secret:    "uoo8h0",
	} 
	
	return new_account
}