// Sample benchmarks to test which function is better for converting
// an integer into a string. First using the fmt.Sprintf function,
// then the strconv.FormatInt function and then strconv.Itoa.
package itoa_test

import (
	"fmt"
	"strconv"
	"testing"
)

/*
 * output:
 * BenchmarkSprintf
 * BenchmarkSprintf-8   	15433765	        77.0 ns/op
 * BenchmarkFormat
 * BenchmarkFormat-8    	326648560	         3.68 ns/op
 * BenchmarkItoa
 * BenchmarkItoa-8      	326164863	         3.71 ns/op
 */

// BenchmarkSprintf provides performance numbers for the
// fmt.Sprintf function.
func BenchmarkSprintf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

// BenchmarkFormat provides performance numbers for the
// strconv.FormatInt function.
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa provides performance numbers for the
// strconv.Itoa function.
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
