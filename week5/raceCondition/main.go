package main

import (
	"fmt"
)

// This program should always return 0, because it is adding and subtracting
// equal numbers. There is no problem once it works sequentially, but concurrent
// work of go routines makes calls different times, even different number of
// times. So, it creates unexpected solution in each compilaiton so, causes
// race condition. Synchronization of go routines is neccesary in this case.

func main() {
	// creating a new integer pointer variable.
	v := new(int)

	*v = 0
	// running the program sequentially.
	runSequentially(v)
	fmt.Println("Running sequentially:", *v)

	*v = 0
	// running the program concurrently.
	runConcurrently(v)
	fmt.Println("Running concurrently:", *v)
}

// add same values to variable v.
func add(v *int) {
	for i := 1; i < 20; i++ {
		*v += i
	}
}

// subtract same values from variable v.
func subtract(v *int) {
	for i := 1; i < 20; i++ {
		*v -= i
	}
}

// running add and subtract functions sequentially. After add function finishes its
// work, subtract function is going to run, so the result is 0 always as expected.
func runSequentially(v *int) {
	for i := 0; i < 256; i++ {
		add(v)
		subtract(v)
	}
}

// running add and subtract functions concurrently will cause a problem called
// a race condition, which creates different results at each compilation.
func runConcurrently(v *int) {
	for i := 0; i < 256; i++ {
		go add(v)
		go subtract(v)
	}
}
