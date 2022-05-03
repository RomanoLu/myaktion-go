package model

type Campaing struct{
	Name string
	DonaitionMinimum float32 
	TargetAmount float32
	Donaitions []Donaition
	Account Account
}