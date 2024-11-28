package repository

import (
	"transactions/entity"
)

type RepositoryAccount interface{
	FindAccountByCpf(string) (bool, entity.Account)
	FindAccountByID(string) (bool, entity.Account)
	ListAccounts()([]entity.Account)
	GetBalance(string) (int)
	UpdateAccount(entity.Account) error
}

type RepositoryTransfer interface{
	CreateTransfer(entity.Transfer) error
	ListTransfers(string)([]entity.Transfer)
}

type Repo struct{

}

