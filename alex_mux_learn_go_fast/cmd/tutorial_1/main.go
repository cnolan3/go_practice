package main

import (
	"fmt"
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
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func arr() {
	// var intArr [3]int32 = [3]int32{1, 2, 3}
	intArr := [3]int32{1, 2, 3} // shorthand initialization

	fmt.Println(intArr)

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

	// for loop over range
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
