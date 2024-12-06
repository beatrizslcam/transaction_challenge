package entity

import "time"

type Transfer struct{
	Id string
	AccountOriginId string
	AccountDestinationId string
	Amount int
	CreatedAt time.Time
}