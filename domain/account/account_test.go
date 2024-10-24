package account_test

import (
	"reflect"
	"testing"
	"transactions/domain/account"
)
 



type MockRepository struct{
	findAccountByCpfFunc func(cpf string) (bool, interface{})
	findAccountByIDFunc  func(id string) (bool, interface{})
	listAccountsFunc func()([]interface{})
	getBalanceFunc func(id string)(balance int)
}

func (m *MockRepository) ListAccounts() []interface{} {
	return m.listAccountsFunc()
}
func (m *MockRepository) FindAccountByCpf(cpf string) (bool, interface{}) {
	return m.findAccountByCpfFunc(cpf)
}

func (m *MockRepository) FindAccountByID(id string) (bool, interface{}) {
	return m.findAccountByIDFunc(id)
}

func  (m *MockRepository) GetBalance(id string) interface{} {
	return m.getBalanceFunc(id)
}

func mockingListAccounts(accounts []account.Account) *MockRepository {
    return &MockRepository{
        listAccountsFunc: func() ([]interface{}) {
			return accounts
	    },
	}
}



func  mockingGetBalance(account account.Account,id string) *MockRepository{
	return &MockRepository{
		getBalanceFunc: func(id string) (balance int){
			return account.Balance
		},
	}
}

func mockingFindByCpf(is_created bool, account account.Account) *MockRepository {
    return &MockRepository{
        findAccountByCpfFunc: func(cpf string) (bool, interface{}) {
            return is_created, account
	    },
	}
}

func mokingFindByID(isCreated bool, account account.Account) *MockRepository{
	return &MockRepository{
		findAccountByIDFunc: func(id string) (bool, interface{}) {
			return isCreated, account
		},
	}
}

 func TestCreateAccount(t *testing.T){
	t.Run("create account", func(t *testing.T){
		t.Parallel()
		
		mockAccount := mockingFindByCpf(false, account.Account{})
		manageAccount := account.ManageAccount{Repo: mockAccount}


		result := manageAccount.CreateAccount("Maria", "12345")
		expect := account.Account{"ugiugiu","Maria", "12345","uoo8h0",0}

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}
	})
	t.Run("Account already exists", func(t *testing.T){
		t.Parallel()
		
		mockAccount := mockingFindByCpf(true, account.Account{"ugiugi","Ana", "17995","uoo8h0",100})
		manageAccount := account.ManageAccount{Repo: mockAccount}

		result := manageAccount.CreateAccount("Ana", "17995")
		expect := account.Account{"ugiugiu","Ana", "17995","uoo8h0",100}

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}
	})
	
 }
 func TestGetAccount(t *testing.T){
	
	t.Run("Get acounnt", func(t *testing.T){
		t.Parallel()

		mockAccount := mokingFindByID(true, account.Account{"ugiugiu","Ana", "17995","uoo8h0",100})
		manageAccount := account.ManageAccount{Repo: mockAccount}

		result := manageAccount.GetAccount("ugiugiu")
		expect := account.Account{"ugiugiu","Ana", "17995","uoo8h0",100}

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}

	})
	t.Run("Account not Found", func(t *testing.T){
		t.Parallel()

		mockAccount := mokingFindByID(true, account.Account{})
		manageAccount := account.ManageAccount{Repo: mockAccount}

		result := manageAccount.GetAccount("ugiugiu")
		expect := account.Account{}

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}

	})
 }
	func TestListAccounts(t *testing.T){

		t.Run("list accounts", func(t *testing.T){
			t.Parallel()

			mockedAccounts := []account.Account{
				 account.Account{"ugiugiu","Ana", "17995","uoo8h0",100},
				 account.Account{"tufuyuy","Maria", "12205","Aoo8h0",300},
			}
			mockRepo := mockingListAccounts(mockedAccounts)
			accountMap := account.ManageAccount{Repo: mockRepo}

			expect := mockedAccounts
			result := accountMap .ListAccounts()
			

			if !reflect.DeepEqual(result,expect){
				t.Errorf("got %v want %v", result, expect)
			}
		})
	}
		
  func TestGetBallance(t *testing.T){
	t.Run("get balance", func(t *testing.T){
		t.Parallel()

		mockAccount := mockingGetBalance(account.Account{"ugiugiu","Ana", "17995","uoo8h0",100},"ugiugiu")
		manageAccount := account.ManageAccount{Repo: mockAccount}

		expect := 100
		result := manageAccount.GetBalance("ugiugiu")

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}
	})
  }