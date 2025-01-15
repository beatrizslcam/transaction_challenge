package account

import (
	"fmt"
	"transactions/entity"
	"transactions/repository"
) 

type ManageAccount struct{
	Repo repository.RepositoryAccount
}


func (ma *ManageAccount) CreateAccount(name string, cpf string)(entity.Account, error){
	existingAccount, err:= ma.Repo.FindAccountByCpf(cpf)
		if err!= nil {
			return existingAccount, fmt.Errorf("error while finding account due to: %v", err)
		}

		if existingAccount.ID != "" {
			return existingAccount, fmt.Errorf("account already exists")
		}	
	newAccount :=entity.Account{
		Name:      name,
	Cpf:       cpf,
	Balance:   0,
	ID:        "ugiugiu",
	Secret:    "uoo8h0",
	} 
	
	return newAccount, nil
}

func (ma *ManageAccount) GetAccount(id string)(entity.Account, error){
	existingAccount, err:= ma.Repo.FindAccountByID(id)
	if err!= nil {
		return existingAccount, fmt.Errorf("error while getting account due to: %v", err)
	}

	if existingAccount.ID == "" {
		return entity.Account{}, fmt.Errorf("account doesn't exists")
	}
		
		return existingAccount, nil
}


func (ma *ManageAccount) ListAccounts() ([]entity.Account, error){
	accountsList,err:= ma.Repo.ListAccounts()
	if err!= nil {
		return nil, fmt.Errorf("error while listing accounts due to: %v", err)
	}
	var result []entity.Account
	result = append(result, accountsList...)
	
	return result,nil
}



func (ma *ManageAccount) GetBalance(id string) (int, error) {	
	existingAccount,err := ma.Repo.FindAccountByID(id)
	if err!= nil {
		return -1, fmt.Errorf("error while get balance due to: %v", err)
	}

		return existingAccount.Balance, nil

}
