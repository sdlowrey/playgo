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
	}()
	return c
}

func main() {
	c := boring("so boring")
	fmt.Println("I'm listening.")
	for i := 0; i < 5; i++ {
		fmt.Printf("You said %s\n", <-c)
	}
	fmt.Println("You're boring. Bye!")
}
