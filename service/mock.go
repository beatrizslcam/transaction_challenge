package service

import "transactions/entity"

type AccountMockRepository struct{
	findAccountByCpfFunc func(string) (bool, entity.Account)
	findAccountByIDFunc  func(string) (bool, entity.Account)
	listAccountsFunc func()([]entity.Account)
	getBalanceFunc func(string)(int)
}

func (m *AccountMockRepository) ListAccounts() []entity.Account {
	return m.listAccountsFunc()
}

func MockingListAccounts(accounts []entity.Account) *AccountMockRepository {
    return &AccountMockRepository{
        listAccountsFunc: func() ([]entity.Account) {
			return accounts
	    },
	}
}

func (m *AccountMockRepository) FindAccountByCpf(cpf string) (bool, entity.Account) {
	return m.findAccountByCpfFunc(cpf)
}

func MockingFindByCpf(is_created bool, account entity.Account) *AccountMockRepository {
    return &AccountMockRepository{
        findAccountByCpfFunc: func(string) (bool, entity.Account) {
            return is_created, account
	    },
	}
}

func (m *AccountMockRepository) FindAccountByID(id string) (bool, entity.Account) {
	return m.findAccountByIDFunc(id)
}

func MockingFindByID(isCreated bool, account entity.Account) *AccountMockRepository{
	return &AccountMockRepository{
		findAccountByIDFunc: func(string) (bool, entity.Account) {
			return isCreated, account
		},
	}
}

func  (m *AccountMockRepository) GetBalance(id string) int{
	return m.getBalanceFunc(id)
}

func  MockingGetBalance(account entity.Account) *AccountMockRepository{
	return &AccountMockRepository{
		getBalanceFunc: func(string) (int){
			return account.Balance
		},
	}
}


type AtuhMock struct{
	generateTokenFunc func(string) string
	validateTokenFunc func(string) bool

}

func (m *AtuhMock) GenerateToken(accountId string) string{
	return m.generateTokenFunc(accountId)
}

func MockingGenerateToken(accountId string) *AtuhMock {
	return &AtuhMock{
		generateTokenFunc: func(string) string {
			return "eyJhbGciOiJIU.eyJzdWIiOiIxMNTE2MjM5MDIyfQ.SflKxwRJS6yJV_adQssw5c"
		},
	}
}

func (m *AtuhMock) ValidateToken(accountId string) bool{
	return m.validateTokenFunc(accountId)
}

func MockingValidateToken(accountId string) *AtuhMock {
	return &AtuhMock{
		validateTokenFunc: func(string) bool {
			return true
		},
	}
}



