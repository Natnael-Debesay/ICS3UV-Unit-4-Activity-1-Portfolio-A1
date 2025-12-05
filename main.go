// Author: Natnael Debesay
// Version: 1.0.0
// Date: 2025-12-05
// Fileoverview: This program counts backward from 100 by 5s.

package main

import (
	"fmt"
)

func main () {
	// print numbers
	for number := 100; number >= 0; number -= 5 {
		fmt.Println(number)
	}

	fmt.Println("\nDone.")
}
