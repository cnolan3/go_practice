package main

import (
	"fmt"
	"strings"
)

func main() {
	var intNum int = 42
	fmt.Println(intNum)

	var floatNum float64 = 12333333.4
	fmt.Println(floatNum)

	printMe("hello world 2")

	arr()

	slice()

	mapTest()

	stringTest()

	structTest()

	pointerTest()
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func arr() {
	// var intArr [3]int32 = [3]int32{1, 2, 3}
	intArr := [3]int32{1, 2, 3} // shorthand initialization

	fmt.Println(intArr)

	// for loop over range of array
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}
}

func slice() {
	// init a slice and append a value to it
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("intSlice: %v\n", intSlice)
	fmt.Printf("the length is %v with capacity %v\n\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("after append one value, intSlice: %v\n", intSlice)
	fmt.Printf("the length is %v with capacity %v\n\n", len(intSlice), cap(intSlice))

	// append another slice to the original slice
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("after append another slice, intSlice: %v\n", intSlice)
	fmt.Printf("the length is %v with capacity %v\n\n", len(intSlice), cap(intSlice))

	// use make to create a slice with length 3 and capacity 8
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("intSlice3: %v\n", intSlice3)
	fmt.Printf("the length is %v with capacity %v\n\n", len(intSlice3), cap(intSlice3))
}

func mapTest() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])

	var age1, ok1 = myMap2["Sarah"]
	var age2, ok2 = myMap2["Jason"]

	if ok1 {
		fmt.Printf("Sarah's age is: %v\n", age1)
	} else {
		fmt.Println("Invalid name: Sarah")
	}

	if ok2 {
		fmt.Printf("Jason's age is: %v\n", age2)
	} else {
		fmt.Println("Invalid name: Jason")
	}

	// for loop over range of map
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	// iterate with index, condition in for keyword
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// iterate with index, condition in break block inside loop
	i = 0
	for {
		if i >= 10 {
			break
		}

		fmt.Println(i)
		i++
	}

	// iterate with index, for loop syntax
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func stringTest() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	/*
		print codes for each character, returns:

		0 114
		1 233
		3 115
		4 117
		5 109
		6 233

		the special character é is two bytes long instead of one,
		this is why the output skips the number 2
	*/
	for i, v := range myString {
		fmt.Println(i, v)
	}

	var myRuneArr = []rune(myString)

	/*
		runes are aliases for int32, and can be used to represent
		individual characters in a string, no matter how many bytes
		the character is.

		iterating over a rune array returns the expected numbers,
		this loop will not skip the number 2:

		0 114
		1 233
		2 115
		3 117
		4 109
		5 233
	*/
	for i, v := range myRuneArr {
		fmt.Println(i, v)
	}

	// concat strings with the + symbol
	var strSlice = []string{"s", "u", "b"}
	var catStr = "aaa"
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v\n", catStr)

	// string builder
	var strSlice2 = []string{"s", "u", "b"}
	var strBuilder strings.Builder
	for i := range strSlice2 {
		strBuilder.WriteString(strSlice2[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Printf("\n%v\n", catStr2)
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
	owner
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("you can make it there!")
	} else {
		fmt.Println("need to fuel up first!")
	}
}

func structTest() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)

	// inline declaration
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{27, 10}

	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	var myEngine3 electricEngine = electricEngine{30, 9, owner{"John"}}

	canMakeIt(myEngine, 50)
	canMakeIt(myEngine3, 50)
}

func pointerTest() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	p = &i
	*p = 1
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)
}
