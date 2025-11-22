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
func Ticker(value string, wg *sync.WaitGroup) {
	defer wg.Done()

	// func that runs every second until 10 secs - and prints the value

	// a timer that runs for two seconds
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Waiting... for ", value)
	<-timer.C
	fmt.Println("Wait done ... for", value)
}

func Monitor() {
	var wg sync.WaitGroup

	// GET all targets
	targets := []string{"google", "facebook", "amazon", "apple"}

	// Create timer for each value
	for _, v := range targets {
		wg.Add(1)
		go Ticker(v, &wg)
	}

	// wg.Add(1)
	// go Ticker("task", &wg)

	wg.Wait()
}
