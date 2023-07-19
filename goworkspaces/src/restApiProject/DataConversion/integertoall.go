package main

import (
	"fmt"
	"strconv"
)

var name = 20 //
var b bool = false
var ST string = "mk"
var Con int64 = 23
var F float32 = 23.4

func Conv(b bool) int8 {
	if b {
		return 1
	}
	return 0
}
func Conversion1() {

	fmt.Println("________________INTEGER To Other Data-Type Conversion______________")
	f := float64(Con)
	fmt.Println("Int to floate", f)
	str := string(Con)
	fmt.Println("integer to string", str)
	rn := rune(Con)
	fmt.Println("int to Rune", rn)
	bt := byte(Con)
	fmt.Println("Int to Byte", bt)
	//	cm:=complex64(Con)//cant convert int to complex64
	//bool(Con) cant convert int to bool
	fmt.Println("_________________FLOATE To Other Data-Type conversion_______________")
	in := int(F)
	fmt.Println("integer to float :", in)
	fmt.Println("float to integer", strconv.Itoa(int(f)))
	fmt.Println(strconv.ParseBool(strconv.FormatInt(Con, 10)))
	fmt.Println(strconv.ParseBool(ST))
	//s, _ :=strconv.Atoi(F)//cant convert float32 to string
	//using strconv we can convert string to other data type only not
	//it converts only string to other and it don't work in other except string
	fmt.Println(strconv.ParseBool("9"))
	r := rune(F)
	//rune means character in go language
	fmt.Println("Float to rune ", r)
	//converting float to byte
	t := byte(F)
	fmt.Println("Float to byte", t)
	c := complex(F, F)
	fmt.Println("Float to complex", c)
	fmt.Println("________________String To other type________________________")

	var S string = "Ok"
	//converting string to boolean
	fmt.Println(strconv.ParseBool(S))
	//converting string ti int
	fmt.Println(strconv.ParseInt(S, 4, 4))
	//converting string to float
	fmt.Println(strconv.ParseFloat(S, 3))
	//rn1:=rune(S)//cant convert string to rune
	//converting string to complex
	fmt.Println(strconv.ParseComplex(S, 3))
	fmt.Println("_____________________Bool to Other type___bool can't be converted to other type")
	//var b bool//we can't convert bool to another any data type
	//a:=int(b)
	//g:=float64(b)
	//s:=string(b)
	//r:=rune(b)
	//c:=complex(b)
	//strconv.ParseInt(b,4,4)
	//bt:=byte(b)
	//converting bool to string
	fmt.Println(strconv.FormatBool(b), "Bool to string")
	fmt.Printf("%T", strconv.FormatBool(b))

	fmt.Println("")
	//fmt.Println(strconv.FormatInt(Con, 1))
	fmt.Println()

}
func main() {
	s := Conv(true)
	fmt.Println(s)
	Conversion1()

}
