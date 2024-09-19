package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	for {
		num := rand.Intn(101) // Generate a random number between 0 and 100
		fmt.Fprintf(conn, "%d\n", num)
		fmt.Println("number: %d", num)
		time.Sleep(1 * time.Second) // Send a number every second
	}
}
