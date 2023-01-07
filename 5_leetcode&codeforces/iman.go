package main

import (
	"fmt"
	"sync"
	"time"
)

// using channel, waitgroup
func main() {
	fmt.Println("ekgemgrmk")
	time1 := time.Now().Second()
	var (
		goroutines = 1000
		numbers    = [1000]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		numbers2   []int
		channel    = make(chan []int, goroutines)
		wg         = sync.WaitGroup{}
	)
	wg.Add(goroutines)

	for i := 1; i <= goroutines; i++ {
		start := len(numbers) * (i - 1) / goroutines
		end := len(numbers) * i / goroutines

		go bye(channel, numbers[start:end], &wg)
	}
	wg.Wait()

	for i := 1; i <= goroutines; i++ {
		select {
		case msg1 := <-channel:
			numbers2 = append(numbers2, msg1...)
		}
	}

	fmt.Println(numbers2)
	fmt.Println(time.Now().Second() - time1)
}

func bye(channel chan []int, numbers []int, wg *sync.WaitGroup) {
	var numbers1 []int
	defer wg.Done()
	for _, num := range numbers {
		time.Sleep(1 * time.Second)
		numbers1 = append(numbers1, num)
	}
	channel <- numbers1
}
// test
