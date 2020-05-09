//creating a race condition using goroutines and rectifying it using sync.Mutex
package main

import (
	"fmt"
	"runtime" //for using NumGoroutine
	"sync"    //for using WaitGroups and Mutex
)

var wg sync.WaitGroup //sync package is used
var m sync.Mutex

func foo() (int, int) {
	gr := 100
	count := 0
	wg.Add(gr)                   //gr no of waitgroups are added
	no := runtime.NumGoroutine() //runtime package is used
	fmt.Println("no of goroutine at start of foo:", no)
	for i := 0; i < gr; i++ {
		go func() {
			m.Lock()          //locking the iterator
			v := count        //iterator
			runtime.Gosched() //this is used to allocate count to other goroutines if they need it
			v++
			count = v
			fmt.Println(count)
			m.Unlock() //unlocking the iterator
			wg.Done()	//signals that a goroutine is complete
		}()
		if i == 50 {
			fmt.Println("no of goroutines at middle of foo:", no)
		}
	}
	fmt.Println("no of goroutines at end of foo:", no)
	wg.Wait() //wait till all goroutines complete
	return count, no
}

func main() {
	n := runtime.NumGoroutine()
	fmt.Println("no of goroutine at start of main:", n)
	s, no := foo() //calling the func foo which creates goroutines
	fmt.Println("no of goroutine after returning to main:", no)
	fmt.Println(s)
}
