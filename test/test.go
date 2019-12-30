package main

import "fmt"

func main(){
	fmt.Println("Hello world")
	fmt.Println("2nd Hello world")
	
	var val1 = 4
	var Val1 = 9
	/*var val2 = 6 + 8
	var val3 = 89*/
	var var3 = 5
	var var4 = 5.34
	
	
	fmt.Println(val1)
	fmt.Println(Val1)
	//fmt.Println(val2)
	fmt.Println(Val1 + val1)
	fmt.Printf("%T\n", var3) //check datatype
	fmt.Printf("%T\n", var4)
}