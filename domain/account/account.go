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

//TODO: armazenar uma secret como hash
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

func (ma *ManageAccount) GetAccount(id string) Account{
	doesExist,existingAccount := ma.Repo.FindAccountByID(id)
		if doesExist {
			return existingAccount.(Account)
		}
		return Account{}
}


func (ma *ManageAccount) ListAccounts() []Account {
	accountsList := ma.Repo.ListAccounts()
	var result []Account
	for _, acc := range accountsList {
		result = append(result, acc.(Account))
	}
	return result
}



func (ma *ManageAccount) GetBalance(id string) int {	
	doesExist,existingAccount := ma.Repo.FindAccountByID(id)
	if doesExist {
		return existingAccount.(Account).Balance
	}
	return -1
}