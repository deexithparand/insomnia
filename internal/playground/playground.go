package playground

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

// global variables
var hostUrl string = "https://pokeapi.co/api/v2/pokemon/"

// // Can read data from Reader
// type Reader interface {
// 	Read([]byte) ([]byte, error)
// }

// // Can write data
// type Writer interface {
// 	Write()
// }

// Reads data from IO Reader
func PlaygroundIoReader() {
	fmt.Println()

	resp, err := http.Get(hostUrl)
	if err != nil {
		panic(err)
	}

	var reader io.Reader = resp.Body
	readBytes, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("Output : ", string(readBytes))

	fmt.Println()
}

// Reads data from Bufio Reader
func PlaygroundBufioReader() {
	fmt.Println()

	resp, err := http.Get(hostUrl)
	if err != nil {
		panic(err)
	}

	bufioReader := bufio.NewReader(resp.Body)
	defer resp.Body.Close()
	for {
		httpLineOp, err := bufioReader.ReadString('\n')

		if len(httpLineOp) > 0 {
			fmt.Println("Line by Line (O/P) : ", httpLineOp)
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("error reading the httpLineOp %v", err)
			break
		}
	}

	fmt.Println()
}

// implements interface
func PlaygroundInterfaceImplementor() {
	fmt.Println()

	fmt.Println()
}
