package main

import (
	"testing"
)

func BenchmarkChannel(b *testing.B) {
	b.StopTimer()
	const worker_count = 1000
	first_input := make(chan int)
	var input = first_input
	for i := 0; i < worker_count; i++ {
		output := make(chan int)
		go worker(input, output)
		input = output
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		first_input <- i
		<-input
	}
}

func worker(input <-chan int, output chan<- int) {
	for val := range input {
		output <- val + 1
	}
}

func TestNoWarning(t *testing.T) {}
