package repository

import (
	"transactions/entity"
)

type Repository interface{
	FindAccountByCpf(string) (bool, entity.Account)
	FindAccountByID(string) (bool, entity.Account)
	ListAccounts()([]entity.Account)
	GetBalance(entity.Account) (int)
}

type Repo struct{

}

