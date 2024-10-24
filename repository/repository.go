package repository



type Repository interface{
	FindAccountByCpf(string) (bool, interface{})
	FindAccountByID(string) (bool, interface{})
	ListAccounts()([]interface{})
	GerBalance(interface{}) (int)
}

type Repo struct{

}

