package core

import "fmt"

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

func Monitor() {

	// GET all targets
	targets := []string{"google", "facebook", "amazon", "apple"}

	fmt.Println("Entered Here")

	TickerMain(targets)

}
