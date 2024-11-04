package account_test

import (
	"reflect"
	"testing"
	"transactions/domain/account"
	"transactions/entity"
)
 



type MockRepository struct{
	findAccountByCpfFunc func(string) (bool, entity.Account)
	findAccountByIDFunc  func(string) (bool, entity.Account)
	listAccountsFunc func()([]entity.Account)
	getBalanceFunc func(string)(int)
}

func (m *MockRepository) ListAccounts() []entity.Account {
	return m.listAccountsFunc()
}
func (m *MockRepository) FindAccountByCpf(cpf string) (bool, entity.Account) {
	return m.findAccountByCpfFunc(cpf)
}

func (m *MockRepository) FindAccountByID(id string) (bool, entity.Account) {
	return m.findAccountByIDFunc(id)
}

func  (m *MockRepository) GetBalance(id string) int{
	return m.getBalanceFunc(id)
}

func mockingListAccounts(accounts []entity.Account) *MockRepository {
    return &MockRepository{
        listAccountsFunc: func() ([]entity.Account) {
			return accounts
	    },
	}
}

func mockingFindByCpf(is_created bool, account entity.Account) *MockRepository {
    return &MockRepository{
        findAccountByCpfFunc: func(string) (bool, entity.Account) {
            return is_created, account
	    },
	}
}

func mokingFindByID(isCreated bool, account entity.Account) *MockRepository{
	return &MockRepository{
		findAccountByIDFunc: func(string) (bool, entity.Account) {
			return isCreated, account
		},
	}
}

func  mockingGetBalance(account entity.Account) *MockRepository{
	return &MockRepository{
		getBalanceFunc: func(string) (int){
			return account.Balance
		},
	}
}

 func TestCreateAccount(t *testing.T){
	t.Run("create account", func(t *testing.T){
		t.Parallel()
		
		mockAccount := mockingFindByCpf(false, entity.Account{})
		manageAccount := account.ManageAccount{Repo: mockAccount}


		result := manageAccount.CreateAccount("Maria", "12345")
		expect := entity.Account{ID: "ugiugiu",Name: "Maria", Cpf: "12345",Secret: "uoo8h0",Balance: 0}

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}
	})
	t.Run("Account already exists", func(t *testing.T){
		t.Parallel()
		
		mockAccount := mockingFindByCpf(true, entity.Account{ID: "ugiugiu",Name: "Ana", Cpf: "17995",Secret: "uoo8h0",Balance: 100})
		manageAccount := account.ManageAccount{Repo: mockAccount}

		result := manageAccount.CreateAccount("Ana", "17995")
		expect := entity.Account{ID: "ugiugiu",Name: "Ana", Cpf: "17995",Secret: "uoo8h0",Balance: 100}

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}
	})
	
 }
 func TestGetAccount(t *testing.T){
	
	t.Run("Get acounnt", func(t *testing.T){
		t.Parallel()

		mockAccount := mokingFindByID(true, entity.Account{ID: "ugiugiu", Name: "Ana", Cpf: "17995", Secret: "uoo8h0", Balance: 100})
		manageAccount := account.ManageAccount{Repo: mockAccount}

		result := manageAccount.GetAccount("ugiugiu")
		expect := entity.Account{ID: "ugiugiu", Name: "Ana", Cpf: "17995", Secret: "uoo8h0", Balance: 100}

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}

	})
	t.Run("Account not Found", func(t *testing.T){
		t.Parallel()

		mockAccount := mokingFindByID(true, entity.Account{})
		manageAccount := account.ManageAccount{Repo: mockAccount}

		result := manageAccount.GetAccount("ugiugiu")
		expect := entity.Account{}

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}

	})
 }
	func TestListAccounts(t *testing.T){

		t.Run("list accounts", func(t *testing.T){
			t.Parallel()

			mockedAccounts := []entity.Account{
				{ID: "ugiugiu", Name: "Ana", Cpf: "17995", Secret: "uoo8h0", Balance: 100},
				{ID: "tufuyuy", Name: "Maria", Cpf: "12205", Secret: "Aoo8h0", Balance: 300},
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

		mockAccount := mockingGetBalance(entity.Account{ID: "tufuyuy", Name: "Maria", Cpf: "12205", Secret: "Aoo8h0", Balance: 300})
		manageAccount := account.ManageAccount{Repo: mockAccount}

		expect := 100
		result := manageAccount.GetBalance("ugiugiu")

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}
	})
  }