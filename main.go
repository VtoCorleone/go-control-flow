package main

import (
	"fmt"
)

func ifStatements() {
	i := 4
	if i == 3 {
		fmt.Println("i is 3")
	} else if i == 2 {
		fmt.Println("i is 2")
	} else {
		fmt.Println("i is not 3 or 2")
	}
}

func switchStatment() {
	j := 1

	switch j {
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("None of those numbers")
	}
}

func basicForLoop() {
	i := 0
	for i < 5 {
		i++
		fmt.Println("i is", i)
	}
}

func commonForLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println("i is", i)
	}
}

func forLoopWithRange() {
	numbers := []int{10, 20, 35, 45}
	// i is the index
	// n is the value of the range/array
	for i, n := range numbers {
		fmt.Println("The index of the loop is", i)
		fmt.Println("The value from the array is", n)
	}
}

func deferExecution() {
	defer fmt.Println("I am run after the function completes")
	fmt.Println("Hello World!")
}

func multipleDefers() {
	defer fmt.Println("I am the first defer statement")
	defer fmt.Println("I am the second defer statement")
	defer fmt.Println("I am the third defer statement")
	fmt.Println("Hello World!")
}

func main() {
	ifStatements()
	switchStatment()
	basicForLoop()
	commonForLoop()
	forLoopWithRange()
	deferExecution()
	multipleDefers()
}
