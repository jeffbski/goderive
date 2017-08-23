// Code generated by goderive DO NOT EDIT.

package pipeline

import (
	"sync"
)

func derivePipeline(f func(lines []string) <-chan string, g func(line string) <-chan int) func([]string) <-chan int {
	return func(a []string) <-chan int {
		b := f(a)
		return deriveJoin(deriveFmap(g, b))
	}
}

func deriveJoin(in <-chan <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		wait := sync.WaitGroup{}
		for c := range in {
			wait.Add(1)
			res := c
			go func() {
				for r := range res {
					out <- r
				}
				wait.Done()
			}()
		}
		wait.Wait()
		close(out)
	}()
	return out
}

func deriveFmap(f func(string) <-chan int, in <-chan string) <-chan (<-chan int) {
	out := make(chan (<-chan int))
	go func() {
		for a := range in {
			b := f(a)
			out <- b
		}
		close(out)
	}()
	return out
}
