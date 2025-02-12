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
		return fmt.Errorf("auth error")
		
	}

	//Get Account Origin
	accountOriginID, err := service.GetAccountIDFromToken(token)
	if err != nil {
		return fmt.Errorf("error getting account id from token")
		
	}

	account, err := mt.RepoAccount.FindAccountByID(accountOriginID)
		if err != nil{
			return fmt.Errorf("error couldn't find account")
		}


	
	// Validate if there is enough balance to transfer
	if payload.Amount > account.Balance{
		return fmt.Errorf("there is not enough amount for that transaction")
		 
	}

	//Get Destination Account
	destinationAccount, err := mt.RepoAccount.FindAccountByID(payload.AccountDestinationId)
		if err != nil{
			return fmt.Errorf("error couldn't find account")
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

	 fmt.Print("Transfer done with success!")

	return nil


}	


func(mt *ManageTransfer) ListTransfers(token string) ([]entity.Transfer, error) {
	//Authenticate
	manageLogin := login.ManageLogin{
		Auth: mt.Auth,
		Repo: mt.RepoAccount, 
	}
	if !manageLogin.IsAuthenticated(token) {
		return nil, fmt.Errorf("auth error")
		
	}

	//Get transfers
	accountID, err := service.GetAccountIDFromToken(token)
	if err != nil {
		return nil, fmt.Errorf("error getting account id from token")
		
	}
	
	transfers,err := mt.RepoTransfer.ListTransfers(accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to list transfers: %w", err)
	}
	
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

