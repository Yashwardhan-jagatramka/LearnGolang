package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Go Routines")
	go func() {
		defer wg.Done()
		for count := 0; count < 1; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 1; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()
	fmt.Println("waiting to finish")
	wg.Wait()
	fmt.Println("\nfinished")
}
