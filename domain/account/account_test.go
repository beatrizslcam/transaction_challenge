package account_test

import (
	"reflect"
	"testing"
	"transactions/domain/account"
	"transactions/entity"
	"transactions/service"
)
 


 func TestCreateAccount(t *testing.T){
	t.Run(
		"create account",
		func(t *testing.T){
		t.Parallel()
		
		mockAccount, _ := service.MockingFindByCpf( entity.Account{})
		manageAccount := account.ManageAccount{Repo: mockAccount}


		result,_ := manageAccount.CreateAccount("Maria", "12345")
		expect := entity.Account{ID: "ugiugiu",Name: "Maria", Cpf: "12345",Secret: "uoo8h0",Balance: 0}

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}
	},
	)
	t.Run("Account already exists", func(t *testing.T){
		t.Parallel()
		
		mockAccount,_ := service.MockingFindByCpf( entity.Account{ID: "ugiugiu",Name: "Ana", Cpf: "17995",Secret: "uoo8h0",Balance: 100})
		manageAccount := account.ManageAccount{Repo: mockAccount}

		result,_ := manageAccount.CreateAccount("Ana", "17995")
		expect := entity.Account{ID: "ugiugiu",Name: "Ana", Cpf: "17995",Secret: "uoo8h0",Balance: 100}

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}
	})
	
 }
 func TestGetAccount(t *testing.T){
	
	t.Run("Get acounnt", func(t *testing.T){
		t.Parallel()

		mockAccount,_ := service.MockingFindByID(entity.Account{ID: "ugiugiu", Name: "Ana", Cpf: "17995", Secret: "uoo8h0", Balance: 100})
		manageAccount := account.ManageAccount{Repo: mockAccount}

		result,_ := manageAccount.GetAccount("ugiugiu")
		expect := entity.Account{ID: "ugiugiu", Name: "Ana", Cpf: "17995", Secret: "uoo8h0", Balance: 100}

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}

	})
	t.Run("Account not Found", func(t *testing.T){
		t.Parallel()

		mockAccount,_ := service.MockingFindByID(entity.Account{})
		manageAccount:= account.ManageAccount{Repo: mockAccount}

		result,_ := manageAccount.GetAccount("ugiugiu")
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
			mockRepo,_ := service.MockingListAccounts(mockedAccounts)
			accountMap := account.ManageAccount{Repo: mockRepo}

			expect := mockedAccounts
			result,_ := accountMap .ListAccounts()
			

			if !reflect.DeepEqual(result,expect){
				t.Errorf("got %v want %v", result, expect)
			}
		})
	}
		
  func TestGetBallance(t *testing.T){
	t.Run("get balance", func(t *testing.T){
		t.Parallel()

		mockAccount,_ := service.MockingGetBalance(entity.Account{ID: "tufuyuy", Name: "Maria", Cpf: "12205", Secret: "Aoo8h0", Balance: 300})
		manageAccount := account.ManageAccount{Repo: mockAccount}

		expect := 100
		result,_ := manageAccount.GetBalance("ugiugiu")

		if !reflect.DeepEqual(result, expect){
			t.Errorf("got %v want %v", result,expect)
		}
	})
  }