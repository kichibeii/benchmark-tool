package benchmark

import (
	"testing"

	"golang.org/x/exp/rand"
)

func exampleFunc(nums []int) {
	for i := range nums {
		nums[i] *= 2
	}
}

// Benchmark function using `BenchmarkFunctions` helper
func BenchmarkExampleFunc(b *testing.B) {
	params := []interface{}{makeRandomNumberSlice()}
	FunctionBenchmark(b, "ExampleFunc", exampleFunc, params)
}

func makeRandomNumberSlice() []int {
	LENGTH := 1000 // Example length of the slice
	numbers := make([]int, LENGTH)
	for i := 0; i < LENGTH; i++ {
		numbers[i] = rand.Intn(1000)
	}
	return numbers
}
