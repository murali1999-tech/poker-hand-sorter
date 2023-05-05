
package main

import (
	"fmt"
	"os"

	"github.com/murali1999-tech/poker-hand-sorter/internal/hand"
	"github.com/murali1999-tech/poker-hand-sorter/pkg/utils"
)

func main() {
	// Read the poker hand from the command line.
	input, err := utils.ReadInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	// Create a new hand from the input.
	h, err := hand.NewHand(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating hand:", err)
		os.Exit(1)
	}

	// Sort the hand and print the result.
	h.Sort()
	fmt.Println(h.String())
}
