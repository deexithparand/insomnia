package core

import (
	"fmt"
	"insomnia/internal/db"
	"sync"
	"time"
)

func Ticker(target db.Target, wg *sync.WaitGroup, stop chan bool, ticker *time.Ticker) {
	defer wg.Done()

	// func that runs every second until 10 secs - and prints the value
	for {
		select {
		case <-stop:
			return
		case t := <-ticker.C:
			fmt.Println(target.Endpoint, " monitored at ", t.Second(), " as per the interval ", target.Interval)
		}
	}
}

func stopAllTickers(tickers [](*time.Ticker), stopChans [](chan bool)) {
	for _, ticker := range tickers {
		ticker.Stop()
	}

	for _, stop := range stopChans {
		stop <- true
	}

	fmt.Println("Closed all the tickers - Nothing to worry")
}

func TickerScheduler(targets []db.Target) {
	var (
		wg        sync.WaitGroup
		tickers   [](*time.Ticker)
		stopChans [](chan bool)
	)

	// Create Timer for each value
	for _, target := range targets {
		wg.Add(1)

		stop := make(chan bool)
		ticker := time.NewTicker(1 * time.Second)

		tickers = append(tickers, ticker)
		stopChans = append(stopChans, stop)

		go Ticker(target, &wg, stop, ticker)
	}

	// Temporarily keeps it running for 10 seconds
	time.Sleep(10 * time.Second)

	defer stopAllTickers(tickers, stopChans)

	wg.Wait()
}
