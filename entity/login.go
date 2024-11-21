package entity


type LoginInterface interface{}

type Login struct{
	CPF string
	Secret string
}