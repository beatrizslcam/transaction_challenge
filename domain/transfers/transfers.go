package transfers

import (
	"time"
	"transactions/repository"
	"transacttions/domain/login"
)

/*/transfers
A entidade Transfer possui os seguintes atributos:

Espera-se as seguintes ações:

GET /transfers - obtém a lista de transferencias da usuaria autenticada.
POST /transfers - faz transferencia de uma Account para outra.
Regras para esta rota

Quem fizer a transferência precisa estar autenticada.
O account_origin_id deve ser obtido no Token enviado.
Caso Account de origem não tenha saldo, retornar um código de erro apropriado
Atualizar o balance das contas*/

type ManageTrasnfer struct{
	Repo repository.Repository
}

type Transfer struct{
	id int
	account_origin_id int
	account_destination_id int
	amount int
	created_at time.Time
}

func(mt *ManageTrasnfer)DoTransfer(token string){
	if login.IsAuthenticated(token) {

	}

	return 


}
//Get Transfers
//DoTransfer
/*
-account origin is authenticated ?
-o account_origin vem de um token enviado
-saldo suficient?
-updatebalance
*/