package model


type Donaition struct{
	Amount float32
	ReceiptRequested bool 
	DonorName string
	Status Status 
	Account Account
}

type Status string
const (
	IN_PROCESS Status ="IN_PROCESS"
	TRANSFERRED Status ="TRANSFERRED"
)