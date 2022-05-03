package model


type Donaition struct{
	Amount float32
	ReceiptRequested bool 
	DonorName string
	Status Status 
	Account Account
}

type Status int
const (
	INPROCESS = iota
	TRANSFERRED 
)