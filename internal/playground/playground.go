package playground

import (
	"fmt"
	"io"
	"net/http"
)

// // Can read data from Reader
// type Reader interface {
// 	Read([]byte) ([]byte, error)
// }

// // Can write data
// type Writer interface {
// 	Write()
// }

func PlaygroundInterface() {
	fmt.Println()

	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/")
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

// interfaces
// type  interface {

// 	Write()
// }
