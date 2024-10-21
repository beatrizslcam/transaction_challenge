package account_test

import (
	"reflect"
	"testing"
	"transactions/domain/account"
)
 



type MockRepository struct{
	FindAccountByCpfFunc func(cpf string) (bool, interface{})
	FindAccountByIDFunc  func(id string) (bool, interface{})
}
func (m *MockRepository) FindAccountByCpf(cpf string) (bool, interface{}) {
	return m.FindAccountByCpfFunc(cpf)
}

func (m *MockRepository) FindAccountByID(id string) (bool, interface{}) {
	return m.FindAccountByIDFunc(id)
}

func mocking(is_created bool, account account.Account) *MockRepository {
	return &MockRepository{
		FindAccountByCpfFunc: func(cpf string) (bool, interface{}) {
			return is_created, account
		},
	}
}


 func TestCreateAccount(t *testing.T){
	t.Run("create account", func(t *testing.T){
		
		mock_account := mocking(false, account.Account{})
		manage_account := account.ManageAccount{Repo: mock_account}


		result := manage_account.CreateAccount("Maria", "12345")
		expect := account.Account{"ugiugiu","Maria", "12345","uoo8h0",0}

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}
	})
	t.Run("Account already exists", func(t *testing.T){
		
		mock_account := mocking(true, account.Account{"ugiugi","Ana", "17995","uoo8h0",100})
		manage_account := account.ManageAccount{Repo: mock_account}

		result := manage_account.CreateAccount("Ana", "17995")
		expect := account.Account{"ugiugiu","Ana", "17995","uoo8h0",100}

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}
	})
	
 }