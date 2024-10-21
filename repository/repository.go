package repository


type Repository interface{
	FindAccountByCpf(string) (bool, interface{})
	FindAccountByID(string) (bool, interface{})
}

type Repo struct{

}

