package main

import (
	"fmt"
	"log"
)

var (
	pass   int
	repass int
)

func main() {
	log.Println("Enter the Password")
	fmt.Scanln(&pass)
	log.Println("Re-Enter Password")
	fmt.Scanln(&repass)
	if pass == repass {
		log.Println("Valide password")
	} else {
		fmt.Println("invalid password")
	}
}
