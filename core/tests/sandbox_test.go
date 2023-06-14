// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"testing"
)

const value float64 = 12345.543211234

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%f", value)
		// if x := fmt.Sprintf("%d", 42); x != "42" {
		// 	b.Fatalf("Unexpected string: %s", x)
		// }
	}
}

func BenchmarkStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// fmt.Sprintf("%f", value)
		strconv.FormatFloat(value, 'f', 6, 64)
		// if x := fmt.Sprintf("%d", 42); x != "42" {
		// 	b.Fatalf("Unexpected string: %s", x)
		// }
	}
}

// func TestResult(t *testing.T) {
// 	fmt.Println(fmt.Sprintf("%f\n", value))
// 	fmt.Println(strconv.FormatFloat(value, 'f', 12, 64))
// 	// require.Equal(t, false, true)
// }
