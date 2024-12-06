package account

import (
	"fmt"
	"transactions/entity"
	"transactions/repository"
) 

type ManageAccount struct{
	Repo repository.RepositoryAccount
}


func (ma *ManageAccount) CreateAccount(name string, cpf string)entity.Account{
	doesExist,existingAccount := ma.Repo.FindAccountByCpf(cpf)
		if doesExist {
			fmt.Printf("Account already exists")
			return existingAccount
		}
	
	newAccount :=entity.Account{
		Name:      name,
	Cpf:       cpf,
	Balance:   0,
	ID:        "ugiugiu",
	Secret:    "uoo8h0",
	} 
	
	return newAccount
}

func (ma *ManageAccount) GetAccount(id string)entity.Account{
	doesExist,existingAccount := ma.Repo.FindAccountByID(id)
		if doesExist {
			return existingAccount
		}
		return entity.Account{}
}


func (ma *ManageAccount) ListAccounts() []entity.Account{
	accountsList := ma.Repo.ListAccounts()
	var result []entity.Account
	result = append(result, accountsList...)
	
	return result
}



func (ma *ManageAccount) GetBalance(id string) int {	
	doesExist,existingAccount := ma.Repo.FindAccountByID(id)
	if doesExist {
		return existingAccount.Balance
	}
	return -1
}
