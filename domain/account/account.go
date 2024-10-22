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
	Secret string // precisa ser armazenado como hash
	Balance  int 
	
}


func (ma *ManageAccount) CreateAccount(name string, cpf string) Account{
	doesExist,existingAccount := ma.Repo.FindAccountByCpf(cpf)
		if doesExist {
			fmt.Printf("Account already exists")
			return existingAccount.(Account)
		}
	
	newAccount :=Account{
		Name:      name,
	Cpf:       cpf,
	Balance:   0,
	ID:        "ugiugiu",
	Secret:    "uoo8h0",
	} 
	
	return newAccount
}