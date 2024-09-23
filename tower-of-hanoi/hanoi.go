package main

import "fmt"

func movePieces(n int, source string, destination string, auxiliary string) {
	if n == 1 {
		fmt.Printf("moving piece 1 from %s to %s\n", source, destination)
		return
	}

	movePieces(n-1, source, auxiliary, destination)
	fmt.Printf("moving piece %d from %s to %s ", n, source, destination)
	movePieces(n-1, auxiliary, destination, source)

}

func main() {
	n := 15
	movePieces(n, "SRC", "DEST", "AUX")
}
