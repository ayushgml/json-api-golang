package main

type User struct {
	ID int
	Valid bool
}

type apiError struct {
	Err string
	Status int
}


