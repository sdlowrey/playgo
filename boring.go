package main

import (
	"fmt"
	"time"
	"math/rand"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	go boring("so boring")
	fmt.Println("I'm listening.")
	time.Sleep(time.Second * 2)
	fmt.Println("You're boring. Bye!")
}