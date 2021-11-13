package calculator

import "sync"

type Calculator struct {
	Input  <-chan int
	Output chan<- int
}

func (c *Calculator) Start() {
	var wg sync.WaitGroup
	workers := 3

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(inputs <-chan int, outputs chan<- int) {
			defer wg.Done()
			for i := range inputs {
				outputs <- i * i
			}
		}(c.Input, c.Output)
	}

	go func() {
		wg.Wait()
		close(c.Output)
	}()
}
