package main

import (
	"fmt"
	"insomnia/insomnia"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	insomnia.Start()

	// keeps the application running - until we terminate
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("App Started - Runs until interrupt")
	<-done
}
