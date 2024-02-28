package main
import (
	"fmt"
	"example.com/calculation"
)

func main() {

	// display_msg := intro.Name("Nabeel Boda")
	// fmt.Println(display_msg)
	// address := intro.Address("Kashmir")
	// fmt.Println("I live in", address)
	
		calculated := calculation.Addition(5,50)
		subtracted := calculation.Subtraction(100,46)
		multiplied := calculation.Multiplication(2,69)
		divided := calculation.Division(500,26850)
		fmt.Println(calculated, subtracted, multiplied, divided)
}