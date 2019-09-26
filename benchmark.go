package json_benchmark

import (
	"testing"
)

// base benchmark tests

func benchmarkMarshalGenericData(result interface{}, f func(interface{}) ([]byte, error), b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(result)
	}
}

func benchmarkUnmarshalGenericData(data []byte, result interface{}, f func([]byte, interface{}) error, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(data, &result)
	}
}

func benchmarkPrettyPrintGenericData(result interface{}, f func(interface{}, string, string) ([]byte, error), b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(result, "", "    ")
	}
}
