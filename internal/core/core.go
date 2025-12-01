package core

// A ticker function that keeps running and prints every second
func Monitor() {

	// Update Unmonitored Endpoints - Needs to run the first time the app starts
	UpdateUnmonitoredEndpoints()

	targets := GetMonitoredEndpoints()

	TickerScheduler(targets)
}
