package main

import "log"

var c = 12.03
var v = 12.12

func Add() {
	x := int(c)
	y := int(v)
	log.Println("Add of Integer", x+y)

	log.Println("Add of Point", (c-float64(x))+(v-float64(y)))
}
func main() {
	Add()
}
