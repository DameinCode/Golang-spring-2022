package main

import "fmt"

func greet(c chan string) {
    fmt.Println("Hello " + <-c + "!")
}

func not_greet(c chan string) {
	fmt.Println("Bye " + <-c + "!")
}


func main() {
	fmt.Println("start main()")

	chanel := make(chan string)
	// chanel <- "John"
	go greet(chanel)
	chanel <- "John"
	go not_greet(chanel)
	chanel <- "Alice"
	
	// time.Sleep(time.Second)
	fmt.Println("end main()")
}