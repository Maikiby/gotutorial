package main 

import (

	"fmt"

	"example.com/greetings"

)

func main() {
	message := greetings.Hello("Niko") 
	fmt.Println(message)
}