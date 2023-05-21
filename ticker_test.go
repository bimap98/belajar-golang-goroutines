package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for data := range ticker.C {
		fmt.Println(data)
	}
}

func TestTick(t *testing.T) {
	tick := time.Tick(1 * time.Second)

	for data := range tick {
		fmt.Println(data)
	}
}
