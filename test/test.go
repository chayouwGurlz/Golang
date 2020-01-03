package main

import (
	"fmt"
	"strconv"
)

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
	
	val1-=1
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
	
	if !(val10 < 3.14 && val10 != 0){
		fmt.Println("The value is correct, bro", val10)
	} else {
		fmt.Println("The value is NOT correct, bro")
	}
	
	var book = "Chirpstory"
	switch book {
		case "History" :
			fmt.Println("The book name is history")
		case "People" :
			fmt.Println("The book name is people")
		case "Technology" :
			fmt.Println("The book name is technology")
		default :
			fmt.Println("We don't know the name!")
	}
	
	for n:=1; n<3; n++{
		fmt.Println("Hello world bro", n)
	}
	
	funcVoid("irine", 4)
	fmt.Println("The return integer is", funcInt(4, 3))
	var res1, res2 = funcMulti(4, 3)
	fmt.Println("The return multiple integer are", res1, res2)
}

func funcVoid(name string, id int){ //void
	fmt.Print("Hello world from func1 " + name)
	fmt.Println(" with id", id +1)
	
	var str = fmt.Sprintf("Hi, this is reply from %s, with value %d", name, id+1)
	fmt.Println(str)
	
	var str2 = strconv.Itoa(id+10)
	fmt.Println("Nice, it will " + str2 + " The Value")
}

func funcInt(numA, numB int) int{ //return single int
	return int(numA + numB)
}

func funcMulti(numA, numB int) (int, int){ //return multiple int
	return numA + numB, numA * numB
}