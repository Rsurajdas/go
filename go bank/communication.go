package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

func communication() {
	fmt.Println("What do you want to do?")
	fmt.Println("Reach us 24/7:-", randomdata.PhoneNumber())
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}
