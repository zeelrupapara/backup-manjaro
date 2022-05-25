package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go test1(&wg)
	wg.Add(1)
	go test2(&wg)
	wg.Add(1)
	go test3(&wg)
	wg.Add(1)
	go test4(&wg)
	wg.Wait()
	fmt.Println("time up:", time.Since((now)))
}
func test1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(200 * time.Millisecond)
	fmt.Println("test1 done")
}
func test2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("test2 done")
}
func test3(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond)
	fmt.Println("test3 done")
}
func test4(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("test4 done")
}
