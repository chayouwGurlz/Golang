package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"
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
	
	var str = strings.ToUpper(" sweet heart") //string func
	fmt.Println("Strings ToUpper " +str)
	str = strings.ToLower(str)
	fmt.Println("Strings ToLower " +str)
	str = strings.TrimSpace(str)
	fmt.Println("Strings TrimSpace " +str)
	str = strings.Trim(str, "art")
	fmt.Println("Strings Trim " +str)
	str = strings.TrimLeft(str, "sw")
	fmt.Println("Strings TrimLeft " +str)
	str = strings.TrimRight(str, " he")
	fmt.Println("Strings TrimRight " +str)
	var strToInt = strings.Count(str, "e")
	fmt.Println("Strings Count", strToInt)
	var strToBool = strings.Contains(str, "swe")
	fmt.Println("Strings Contains", strToBool)
	strToBool = strings.ContainsAny(str, "swe")
	fmt.Println("Strings ContainsAny", strToBool)
	strToInt = strings.Index(str, "e")
	fmt.Println("Strings Index", strToInt)
	strToInt = strings.LastIndex(str, "e")
	fmt.Println("Strings LastIndex", strToInt)
	strToInt = strings.LastIndexAny(str, "e")
	fmt.Println("Strings LastIndexAny", strToInt)
	str = strings.Replace(str, "e", "*", -1)
	fmt.Println("Strings Replace", str)
	str = "sweet.heart."
	var arrStr = strings.Split(str, ".")
	fmt.Println("Strings Split", arrStr)
	arrStr = strings.SplitAfter(str, ".")
	fmt.Println("Strings SplitAfter", arrStr)
	str = strings.Join([]string{"Hello", str, "<3"}, ".")
	fmt.Println("Strings Join", str)
	str = strings.Repeat(str, 2)
	fmt.Println("Strings Repeat", str)
	
	var val11 = math.Ceil(0.35) //math function
	fmt.Println("Math Ceil", val11)
	val11 = math.Max(0.35, -7.90)
	fmt.Println("Math Max", val11)
	val11 = math.Min(0.35, -7.90)
	fmt.Println("Math Min", val11)
	val11 = math.Remainder(0.35, -7.90)
	fmt.Println("Math Remainder", val11)
	val11 = math.Sqrt(-7.90)
	fmt.Println("Math Sqrt", val11)
	val11 = math.Mod(0.35, -7.90)
	fmt.Println("Math Mod", val11)
	
	var arrInt = [...]int{8,9}
	var arrString = [...]string{"Dewi","Sartika"}
	var arrString1 = [...]string{"Kapten", "Pattimura"}
	fmt.Println("Array of Int", arrInt)
	fmt.Println("Array of String", arrString)
	fmt.Println("Array of String1", arrString1)
	fmt.Println("Sum Array of Int", arrInt[0] + arrInt[1])
	fmt.Println("Comparation arrString & arrString1", arrString == arrString1)
	
	var multiArrayInt = [...][3]int{{7,8,9}, {1,2}}
	fmt.Println("Multiple Array of Int", multiArrayInt)
	fmt.Println("Min Multiple Array of Int row 3", multiArrayInt[0][2] - multiArrayInt[1][2])
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