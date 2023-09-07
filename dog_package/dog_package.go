// Package Dog_package package converts humans age to dog age
package dog_package

import "fmt"

// Years collect humans age (h_age) and coverts it to an equivalent of the age in dog years.
//I this instace, we assume that 1 human year = 7 dog years. So, we save the value "7" in a d_age container, as an int
// cal holds the value of the mathematical calculation
// Print the calculated dog years equivalent.
// Return the calculated dog years equivalent.
func Years(h_age int) int {
	d_age := 7
	cal := h_age * d_age
	fmt.Printf("The dog years equivalent of the age is %v", cal)
	return cal
}
