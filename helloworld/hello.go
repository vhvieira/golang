// for docs use
// godoc -http :8000
// http://localhost:8000/pkg/
package main

import "fmt"

func Hello() string {
	return "Hello, world, from Victor"
}

func main() {
	fmt.Println(Hello())
}
