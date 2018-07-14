package main

type iataValue struct {
	Iata string
	Name string
}

type value struct {
	Amount float64
	Currency string
}

type descriptionValue struct {
	Code string
	Description string
}