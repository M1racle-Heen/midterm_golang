package main

import "sync"

var (
	stores map[string]string
	wg     sync.WaitGroup
	m      sync.Mutex
)

func main() {
	stores = map[string]string{"1": "soda", "2": "bread", "3": "bread"}
	wg.Add(1)
	go handleRequests()
	wg.Wait()
}
