package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	port := 3000
	// startWebServer --> function as an object
	_, err := startWebServer(port, 2) //function invocation, "_" ignore the first result
	fmt.Println(err)
	fmt.Println("Server started on port", port)
}

func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")
	//do things
	fmt.Println("Number of retries", numberOfRetries)
	return port, nil
}
