package core

import (
	"fmt"
	"sync"
	"time"
)

func Ticker(value string, wg *sync.WaitGroup, done chan bool, ticker *time.Ticker) {
	defer wg.Done()

	// func that runs every second until 10 secs - and prints the value
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			fmt.Println(value, " ticks at ", t.Second())
		}
	}
}

func TickerScheduler(targets []string) {
	var (
		wg         sync.WaitGroup
		tickerList [](*time.Ticker)
		doneList   [](chan bool)
	)

	// Create Timer for each value
	for _, v := range targets {
		wg.Add(1)

		done := make(chan bool)
		ticker := time.NewTicker(1 * time.Second)

		tickerList = append(tickerList, ticker)
		doneList = append(doneList, done)

		go Ticker(v, &wg, done, ticker)
	}

	// Temporarily keeps it running for 10 seconds
	time.Sleep(10 * time.Second)

	for _, ticker := range tickerList {
		ticker.Stop()
	}

	for _, done := range doneList {
		done <- true
	}

	wg.Wait()
}
