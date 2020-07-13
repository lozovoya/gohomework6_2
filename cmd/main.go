package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

type Transaction struct {
	Amount int64
	Moment time.Time
}

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

	const amountOfTransactions = 1_000_000_0
	const max = 1_000_000

	transactions := make([]Transaction, amountOfTransactions)

	for index := range transactions {
		transactions[index].Amount = int64(rand.Intn(max))
		//transactions[index].Amount = 1
		rand.Seed(int64(time.Now().Nanosecond()))
		transactions[index].Moment = time.Date(2019, time.Month(rand.Intn(11)), rand.Intn(30), rand.Intn(23), rand.Intn(59), rand.Intn(59), 0, time.UTC)
	}

	var m = make(map[time.Month][]*Transaction)

	for i := range transactions {
		m[transactions[i].Moment.Month()] = append(m[transactions[i].Moment.Month()], &transactions[i])
	}

	wg := sync.WaitGroup{}
	wg.Add(len(m))

	for month := time.Month(1); month <= time.Month(12); month++ {
		//fmt.Println(month, m[month])
		k := month
		if m[month] != nil {
			go func() {
				sum := Sum(m[k])
				fmt.Println(k, sum)
				wg.Done()
			}()
		}
	}
	wg.Wait()

}

func Sum(transactions []*Transaction) int64 {
	result := int64(0)
	for _, transaction := range transactions {
		result += transaction.Amount
	}
	return result
}
