package main

import (
	"fmt"
	"insomnia/insomnia"
	"insomnia/internal"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Config Scripts
	configFilePath := "./config.test.yml"
	config := insomnia.Config(configFilePath)

	// DB Scripts
	insomnia.DB(config)
	defer internal.Close()

	insomnia.Start()

	// keeps the application running - until we terminate
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("App Started - Runs until interrupt")
	<-done
}
