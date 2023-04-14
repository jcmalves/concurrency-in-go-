package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "Hello, world"

	var mutext sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, universe!", &mutext)
	go updateMessage("Hello, cosmos!", &mutext)
	wg.Wait()

	fmt.Println(msg)
}
