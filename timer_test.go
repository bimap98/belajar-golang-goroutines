package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println(time.Now())

	data := <-timer.C
	fmt.Println(data)
}

func TestAfter(t *testing.T) {
	timer := time.After(4 * time.Second)
	fmt.Println(time.Now())
	data := <-timer
	fmt.Println(data)

}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(3*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())

	group.Wait()
}
