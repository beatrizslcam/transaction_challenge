package service

import (
	"fmt"
	"transactions/entity"
)

type AccountMockRepository struct{
	findAccountByCpfFunc func(string) (entity.Account,error)
	findAccountByIDFunc  func(string) (entity.Account,error)
	listAccountsFunc func()([]entity.Account,error)
	getBalanceFunc func(string)(int,error)
	updateAccountFunc func(entity.Account) error
}

func (m *AccountMockRepository) ListAccounts() ([]entity.Account,error) {
	return m.listAccountsFunc()
}

func MockingListAccounts(accounts []entity.Account) (*AccountMockRepository, error) {
    return &AccountMockRepository{
        listAccountsFunc: func() ([]entity.Account, error) {
			return accounts, nil
	    },
	},nil
}

func (m *AccountMockRepository) FindAccountByCpf(cpf string) (entity.Account,error) {
	return m.findAccountByCpfFunc(cpf)
}

func MockingFindByCpf( account entity.Account) (*AccountMockRepository, error) {
    return &AccountMockRepository{
        findAccountByCpfFunc: func(string) ( entity.Account, error) {
            return account,nil
	}, 
},nil
		}
func (m *AccountMockRepository) FindAccountByID(id string) (entity.Account,error) {
	return m.findAccountByIDFunc(id)
}


func  (m *AccountMockRepository) GetBalance(id string) (int, error){
	return m.getBalanceFunc(id)
}

func MockingFindByID( account entity.Account) (*AccountMockRepository, error){
	return &AccountMockRepository{
		findAccountByIDFunc: func(string) (entity.Account, error) {
			return account, nil
		},
	}, nil
}
func  MockingGetBalance(account entity.Account) (*AccountMockRepository, error){
	return &AccountMockRepository{
		getBalanceFunc: func(string) (int, error){
			return account.Balance, nil
		},
	}, nil
}


func  (m *AccountMockRepository) UpdateAccount( account entity.Account) error{
	return m.updateAccountFunc(account)
}

func  UpdateAccount(account entity.Account) (*AccountMockRepository, error){
	return &AccountMockRepository{
		updateAccountFunc: func(entity.Account) error{
			return nil
		},
	}, nil
}


type AuthMock struct{
	generateTokenFunc func(string) (string, error)
	validateTokenFunc func(string) bool

}

func (m *AuthMock) GenerateToken(accountId string) (string, error){
	fmt.Println("Mock GenerateToken called")
	return m.generateTokenFunc(accountId)
}

func (m *AuthMock) ValidateToken(accountId string) bool{
	return m.validateTokenFunc(accountId)
}


func MockingGenerateToken(generateTokenFunc func(string)(string, error)) *AuthMock {
	return &AuthMock{
		generateTokenFunc: generateTokenFunc,
	}
}
func MockingValidateToken(validateTokenFunc func(string) bool) *AuthMock {
	return &AuthMock{
		validateTokenFunc: validateTokenFunc,
	}
}



