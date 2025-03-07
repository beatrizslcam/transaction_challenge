package repository

import (
	"context"
	"fmt"
	"transactions/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)
type TransferRepo struct {
	Pool *pgxpool.Pool
}


func (r *TransferRepo) CreateTransfer(cpf string) (entity.Transfer, error) {
	query := "SELECT * FROM accounts WHERE cpf = $1 "
	row := r.Pool.QueryRow(context.Background(), query, cpf)

	var transfer entity.Transfer
	err := row.Scan(&transfer.ID, &transfer.AccountOriginId, &transfer.AccountDestinationId, &transfer.Amount, &transfer.CreatedAt)
	if err != nil {
		fmt.Errorf("failed get account  due to: %w", err)
		return entity.Transfer{}, err
	}

	return transfer, nil
}


func (r *TransferRepo) ListTransfers(account_id string) ([]entity.Transfer, error) {
	query := "SELECT * FROM transfer where account_id = $1"
	rows, err := r.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transfers []entity.Transfer
	for rows.Next() {
		var transfer entity.Transfer
		err := rows.Scan(&transfer.ID, &transfer.AccountOriginId, &transfer.AccountDestinationId, &transfer.Amount, &transfer.CreatedAt)
		if err != nil {
			fmt.Errorf("failed to list transfers due to: %w", err)
			return nil, err
		}
		transfers = append(transfers, transfer)
	}
	if err = rows.Err(); err != nil {
		fmt.Errorf("failed to list Transfer due to: %w", err)
		return nil, err
	}

	return transfers, nil
}

