package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, length int) chan string {
	c := make(chan string)
	rand.Seed(rand.Int63())
	go func() {
		for i := 0; i < length; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func main() {
	fmt.Println("I'm listening.")
	for whatever := range boring("something dull", 5) {
		fmt.Printf("You said %s\n", whatever)
	}
	fmt.Println("You're boring. Bye!")
}
