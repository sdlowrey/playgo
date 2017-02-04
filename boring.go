package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) chan string {
	c := make(chan string)
	rand.Seed(rand.Int63())
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
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
