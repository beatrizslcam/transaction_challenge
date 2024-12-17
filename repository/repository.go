package repository

import (
	"transactions/entity"
)

type RepositoryAccount interface{
	FindAccountByCpf(string) (entity.Account,error)
	FindAccountByID(string) (entity.Account, error)
	ListAccounts()([]entity.Account, error)
	GetBalance(string) (int,error)
	UpdateAccount(entity.Account) error
}

type RepositoryTransfer interface{
	CreateTransfer(entity.Transfer) error
	ListTransfers(string)([]entity.Transfer, error)
}

type Repo struct{

}

