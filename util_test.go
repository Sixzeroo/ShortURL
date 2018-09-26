package main

import (
	"testing"
	"fmt"
	"sync"
)

func TestIdGenerator(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			tem := IdGenerator(1)
			fmt.Println(tem)
		}()
	}
	wg.Wait()
}