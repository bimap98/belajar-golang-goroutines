package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(2 * time.Second)
			group.Done()
		}()
	}

	fmt.Println(runtime.NumCPU())
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThread)
	fmt.Println(runtime.NumGoroutine())

	group.Wait()

}

func TestChangeThread(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(2 * time.Second)
			group.Done()
		}()
	}

	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(6)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThread)
	fmt.Println(runtime.NumGoroutine())

	group.Wait()

}
