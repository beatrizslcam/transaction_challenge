package repository

import (
	"transactions/entity"
)

type Repository interface{
	FindAccountByCpf(string) (bool, entity.Account)
	FindAccountByID(string) (bool, entity.Account)
	ListAccounts()([]entity.Account)
	GetBalance(string) (int)
}

type Repo struct{

}

