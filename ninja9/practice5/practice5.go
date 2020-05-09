//creating a race condition using goroutines and using sync/atomic package to overcome it
package main

import (
	"fmt"
	"runtime"     //for using NumGoroutine
	"sync"        //for using WaitGroups
	"sync/atomic" //for using atomic package's addint64 and loadint64
)

var wg sync.WaitGroup //sync package is used
var count int64

func foo() (int64, int) {
	gr := 100
	wg.Add(gr)                   //gr no of waitgroups are added
	no := runtime.NumGoroutine() //runtime package is used
	fmt.Println("no of goroutine at start of foo:", no)
	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			fmt.Println(atomic.LoadInt64(&count))
			wg.Done() //signals that a goroutine is complete
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
