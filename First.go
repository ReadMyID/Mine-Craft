package main

import "fmt"

type Person struct {
	Address Address
}

type Address struct {
	DoorNbr    int
	StreetName string
	City       string
}
func Address tostring() string{
	
}
func showAddress(p Person) {
	fmt.Println(p.Address.tostring())
}
