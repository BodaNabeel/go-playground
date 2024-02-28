package main 
import (
	"fmt"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Nabeel Boda")
	fmt.Println(message)
}