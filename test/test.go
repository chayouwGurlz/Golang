package main

import "fmt"

func main(){
	fmt.Println("Hello world")
	fmt.Println("2nd Hello world")
	
	var val1 = 4
	var Val1 = 9
	/*var val2 = 6 + 8 //comment style
	var val3 = 89*/
	var val3 = 5
	var val5 int = 6
	var val2 int
	var val4 = 5.34
	var val6 float64 = 5.89
	var val7 float64
	var val8, val9, val10 = "foo", 8, 3.14
	
	val1--
	fmt.Println(val1)
	val1++
	fmt.Println(Val1)
	fmt.Println(Val1 % val5)
	fmt.Println(Val1 + val1) //arithmetic
	fmt.Println(val3 * val1)
	
	//fmt.Println(val5 - val4) //(mismatched types int and float64)
	//fmt.Println(val2 / val7) //(mismatched types int and float64)
	//fmt.Println(val6 % val4) //(operator % not defined on float64)
	
	fmt.Println(val2)
	fmt.Println(val7)
	fmt.Println(val8)
	fmt.Println(val9)
	fmt.Println(val10)
	
	fmt.Printf("%T\n", val3) //check datatype
	fmt.Printf("%T\n", val5)
	fmt.Printf("%T\n", val4)
	fmt.Printf("%T\n", val6)
}