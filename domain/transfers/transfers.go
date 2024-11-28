package transfers

import (
	"fmt"
	"time"
	"transactions/domain/login"
	"transactions/entity"
	"transactions/repository"
	"transactions/service"

	"github.com/google/uuid"
)

/*/transfers
A entidade Transfer possui os seguintes atributos:

Espera-se as seguintes ações:

GET /transfers - obtém a lista de transferencias da usuaria autenticada.
POST /transfers - faz transferencia de uma Account para outra.
Regras para esta rota



Caso Account de origem não tenha saldo, retornar um código de erro apropriado
Atualizar o balance das contas*/

type ManageTransfer struct{
	RepoTransfer repository.RepositoryTransfer
	RepoAccount repository.RepositoryAccount
	Auth service.Auth
}

type TransferPayload struct{
	AccountDestinationId string
	Amount int
}

func(mt *ManageTransfer)DoTransfer(token string,payload TransferPayload) error{

	//Authenticate
	manageLogin := login.ManageLogin{
		Auth: mt.Auth,
		Repo: mt.RepoAccount, 
	}
	if !manageLogin.IsAuthenticated(token) {
		return fmt.Errorf("Auth error")
		
	}

	//Get Account Origin
	accountOriginID, err := service.GetAccountIDFromToken(token)
	if err != nil {
		return fmt.Errorf("Error getting account id from token")
		
	}

	exists,account := mt.RepoAccount.FindAccountByID(accountOriginID)
		if exists == false {
			return fmt.Errorf("Error couldn't find account")
		}


	
	// Validate if there is enough balance to transfer
	if payload.Amount > account.Balance{
		return fmt.Errorf("There is not enough amount for that transaction")
		 
	}

	//Get Destination Account
	exists,destinationAccount := mt.RepoAccount.FindAccountByID(payload.AccountDestinationId)
		if exists == false {
			return fmt.Errorf("Error couldn't find account")
		}

		err = mt.updateBalance(account, destinationAccount, payload.Amount)
		if err != nil {
			return err
		}

	//Create Transfer
	transfer := entity.Transfer{
		Id:                  uuid.New().String(),
		AccountOriginId:     accountOriginID,
		AccountDestinationId: payload.AccountDestinationId,
		Amount:              payload.Amount,
		CreatedAt:           time.Now(),
	}

	 err = mt.RepoTransfer.CreateTransfer(transfer)
	 if err != nil {
		return fmt.Errorf("failed to create Transfer due to: %w", err)
	 }

	return nil


}	


func(mt *ManageTransfer) ListTrasnfers(token string) ([]entity.Transfer, error) {
	//Authenticate
	manageLogin := login.ManageLogin{
		Auth: mt.Auth,
		Repo: mt.RepoAccount, 
	}
	if !manageLogin.IsAuthenticated(token) {
		return nil, fmt.Errorf("Auth error")
		
	}


	//Get transfers
	accountID, err := service.GetAccountIDFromToken(token)
	if err != nil {
		return nil, fmt.Errorf("Error getting account id from token")
		
	}
	
	transfers := mt.RepoTransfer.ListTransfers(accountID)
	return transfers, nil
}


func (mt *ManageTransfer) updateBalance(accountOrigin entity.Account, accountDestination entity.Account, amount int) error {
	accountOrigin.Balance -= amount
	accountDestination.Balance += amount

	if err := mt.RepoAccount.UpdateAccount(accountOrigin); err != nil {
		return  fmt.Errorf("failed to update origin account: %w", err)
	}
	if err := mt.RepoAccount.UpdateAccount(accountDestination); err != nil {
		return fmt.Errorf("failed to update destination account: %w", err)
	}

    fmt.Printf("Balance updated with success!") 

	return nil
}

