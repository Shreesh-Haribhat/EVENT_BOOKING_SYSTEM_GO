package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello There Well come to my applcaition")

	done := make(chan bool)

	go doSomething("fast 1", done)
	go doSomething("fast 2", done)
	go didSomething("slowwwwwwwww", done)
	go doSomething("fast 3", done)

	<-done
	<-done
	<-done
	<-done
}

func doSomething(str string, done chan bool) {
	fmt.Println(str)
	done <- true
}

func didSomething(str string, done chan bool) {
	time.Sleep(3 * 100)
	fmt.Println(str)
	done <- true
}
