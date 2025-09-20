package main

import (
	"fmt"
	"sync"
)

var (
	m     map[int]int
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	numWorkers := 500
	m = make(map[int]int)
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			mutex.Lock()
			m[id] = m[id] + 1
			mutex.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("Map length = %d\n", len(m)) // should always be 500
	fmt.Println("Sample values:", m)

}
