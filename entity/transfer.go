package entity

import "time"

type Transfer struct{
	ID string
	AccountOriginId string
	AccountDestinationId string
	Amount int
	CreatedAt time.Time
}