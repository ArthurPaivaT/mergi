package main

import (
	"fmt"

	"github.com/ArthurPaivaT/mergi/server"
)

func main() {

	serverErrChan := make(chan error)

	go server.Start(serverErrChan)

	serverErr := <-serverErrChan
	fmt.Println(fmt.Errorf("Error Serving: %w", serverErr))

	return
}
