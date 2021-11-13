package acmp_concurrent

import (
	"github.com/erkkke/onetech_internship_test/acmp"
	"sync"
)

func Difficulties(urls []string) map[string]float64 {
	var wg sync.WaitGroup
	var mu sync.RWMutex
	sm := make(chan struct{}, 5)
	res := make(map[string]float64)

	for _, url := range urls {
		wg.Add(1)
		sm <- struct{}{}
		go worker(url, &mu, &wg, sm, &res)
	}

	wg.Wait()
	close(sm)

	return res
}

func worker(url string, mu *sync.RWMutex, wg *sync.WaitGroup, sm <-chan struct{}, res *map[string]float64) {
	defer func() { <-sm }()
	defer wg.Done()

	difficulty := acmp.Difficulty(url)

	mu.Lock()
	defer mu.Unlock()
	(*res)[url] = difficulty
}
