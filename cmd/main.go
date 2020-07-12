package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

func main() {

	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	type transaction struct {
		Amount int64
		Moment time.Time
	}

	const amountOfTransactions = 10
	const max = 1_000_000
	const partsCount = 10

	transactions := make([]transaction, amountOfTransactions)

	for index := range transactions {
		transactions[index].Amount = int64(rand.Intn(max))
		rand.Seed(int64(time.Now().Nanosecond()))
		transactions[index].Moment = time.Date(2020, time.Month(rand.Intn(11)+1), rand.Intn(29)+1, rand.Intn(22)+1, rand.Intn(58)+1, rand.Intn(58)+1, 0, time.UTC)
	}

	fmt.Println(transactions)
}

//func sum (transactions []int64) int64 {
//	result := int64(0)
//	for _, transaction := range transactions {
//		result += transaction
//	}
//	return result
//}
//
//func SumConcurrently (transactions []int64, goroutines int) int64 {
//	wg := sync.WaitGroup{}
//	wg.Add(goroutines)
//	total := int64(0)
//	partSize := len(transactions)/goroutines
//	for i := 0; i < goroutines; i++ {
//		part := transactions[i*partSize : (i+1)* partSize]
//		go func() {
//			fmt.Println("start")
//			total += sum(part)
//			wg.Done()
//		}()
//	}
//
//	wg.Wait()
//	return total
//}
