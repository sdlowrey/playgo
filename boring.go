package main

import (
	"fmt"
	"time"
	"math/rand"
)

func boring(msg string, c chan string) {
	rand.Seed(rand.Int63())
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go boring("so boring", c)
	fmt.Println("I'm listening.")
	for i := 0; i < 5; i++ {
		fmt.Printf("You said %s\n", <-c)
	}
	fmt.Println("You're boring. Bye!")
}