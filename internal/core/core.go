package core

// core - code
/**

1. Fetch all targets from the database, including id, URL, interval, last_run, and next_run.
2. Start a parallel goroutine for each target with its own ticker based on the interval.
3. Each ticker wakes up periodically and checks if the current time matches or exceeds next_run in the DB.
4. If it’s time to run, hit the endpoint and record success/failure metrics.
5. Update last_run to now and compute + update next_run (now + interval).
6. Continue running indefinitely unless the app stops or the target is deleted.
7. On app shutdown, stop all tickers and signal all goroutines to exit.

*/

// A ticker function that keeps running and prints every second
// NOTE : We also need to check if the time here equates the one in the DB

func Monitor() {

	// Update Unmonitored Endpoints - Needs to run the first time the app starts
	UpdateUnmonitoredEndpoints()

	targets := GetMonitoredEndpoints()

	TickerScheduler(targets)
}
