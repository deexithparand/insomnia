package internal

import (
	"fmt"
	"sync"
	"time"
)

// core - code
/**
Get the target ids - url - interval - Last Run is Empty
Parallel Ticker Function that keeps running with a target each
when ticket.timer >= next_run - update last_run & hit the endpoint
update stats in the db about the endpoint & next_run
again the ticker keeps going until the app stops or the endpoint deleted
*/

// A ticker function that keeps running and prints every second
// NOTE : We also need to check if the time here equates the one in the DB
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

func Monitor() {
	var (
		wg         sync.WaitGroup
		tickerList [](*time.Ticker)
		doneList   [](chan bool)
	)

	// GET all targets
	targets := []string{"google", "facebook", "amazon", "apple"}

	// Create timer for each value
	for _, v := range targets {
		wg.Add(1)

		done := make(chan bool)
		ticker := time.NewTicker(1 * time.Second)

		tickerList = append(tickerList, ticker)
		doneList = append(doneList, done)

		go Ticker(v, &wg, done, ticker)
	}

	time.Sleep(10 * time.Second)

	for _, ticker := range tickerList {
		ticker.Stop()
	}

	for _, done := range doneList {
		done <- true
	}

	wg.Wait()
}
