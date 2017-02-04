package main

import (
	"fmt"
)

func boring(msg string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
		close(c)
	}()
	return c
}

func fanIn(in1, in2 <-chan string) chan string{
	c := make(chan string)
	go func() { for { c <- <-in1}}()
	go func() { for { c <- <-in2}}()
	return c
}

func main() {
	c := fanIn(boring("blah blah"), boring("shut up"))
	for  i := 0; i < 10; i++{
		fmt.Println(<- c)
	}
	fmt.Println("You're boring. Bye!")
}
