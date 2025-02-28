package repositories

import (
	"context"
	"fmt"
	"transactions/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)
type AccountRepo struct {
	Pool *pgxpool.Pool
}

func (r *AccountRepo) FindAccountByCpf(cpf string) (entity.Account, error) {
	query := "SELECT * FROM accounts WHERE cpf = $1 "
	row := r.Pool.QueryRow(context.Background(), query, cpf)

	var account entity.Account
	err := row.Scan(&account.ID, &account.Name, &account.Cpf, &account.Secret, &account.Balance)
	if err != nil {
		fmt.Errorf("failed get account  due to: %w", err)
		return entity.Account{}, err
	}

	return account, nil
}


func (r *AccountRepo) FindAccountByID(id string) (entity.Account, error) {
	query := "SELECT * FROM accounts WHERE ID = $1"
	row := r.Pool.QueryRow(context.Background(), query, id)

	var account entity.Account
	err := row.Scan(&account.ID, &account.Name, &account.Cpf, &account.Secret, &account.Balance)
	if err != nil {
		fmt.Errorf("failed get account  due to: %w", err)
		return entity.Account{}, err
	}

	return account, nil
}

func (r *AccountRepo) ListAccounts() ([]entity.Account, error) {
	query := "SELECT * FROM accounts"
	rows, err := r.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []entity.Account
	for rows.Next() {
		var account entity.Account
		err := rows.Scan(&account.ID, &account.Name, &account.Cpf, &account.Secret, &account.Balance)
		if err != nil {
			fmt.Errorf("failed to list accounts due to: %w", err)
			return nil, err
		}
		accounts = append(accounts, account)
	}
	if err = rows.Err(); err != nil {
		fmt.Errorf("failed to list accounts due to: %w", err)
		return nil, err
	}

	return accounts, nil
}

func (r *Repo) UpdateAccount(id string, balance int) error {
	query := "Update accounts set balance = $1 where id = $2"
	_, err := r.Pool.Exec(context.Background(), query, id)
	if err != nil {
		fmt.Errorf("failed to update account: %s, due to: %w",id, err)
		return err
	}
	fmt.Printf("Account %s updated with success!", id)

	return nil
}

