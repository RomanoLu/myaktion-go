package model

type Campaign struct{
	ID uint
	Name string
	DonaitionMinimum float32 
	TargetAmount float32
	Donaitions []Donaition
	Account Account
}