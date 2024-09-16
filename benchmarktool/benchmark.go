package benchmarktool

import (
	"reflect"
	"testing"
)

// BenchmarkFunctions is a utility function to benchmark arbitrary functions with parameters.
func FunctionBenchmark(b *testing.B, funcName string, fn interface{}, params []interface{}) {
	b.Run(funcName, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()

			// Prepare arguments for the function using reflection
			args := make([]reflect.Value, len(params))
			for i, param := range params {
				args[i] = reflect.ValueOf(param)
			}

			b.StartTimer()

			// Call the function using reflection
			reflect.ValueOf(fn).Call(args)
		}
	})
}
