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

func main() {
	donnie := boring("blah blah")
	walter := boring("shut the fuck up")
	for  i := 0; i < 5; i++{
		fmt.Println(<- donnie)
		fmt.Println(<- walter)
	}
	fmt.Println("You're boring. Bye!")
}
