package main

import (
	"fmt"

	"github.com/limjunseok0929/learngo/export"
)

/* Function */
func someFunction() {
	fmt.Println("This is Function!!")
}

func main() {
	/* Print */
	fmt.Println("Hello World")

	/* Input */
	var a int
	fmt.Scanln(&a)
	fmt.Println(a)

	/* Var */
	name := "Junseok Lim" // = var name string = "Junseok Lim"
	fmt.Println(name)

	/* If, Else */
	var number int
	fmt.Scanln(&number)
	if number > 10 {
		fmt.Printf("%d is bigger than 10\n", number)
	} else {
		fmt.Printf("%d is not bigger than 10\n", number)
	}

	/* For */
	var star int
	fmt.Scanln(&star)
	for i := 1; i <= star; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	/* Function */
	someFunction()

	/* Import */
	export.Export()
}
