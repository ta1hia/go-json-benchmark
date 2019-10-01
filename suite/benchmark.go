package suite

import (
	"testing"
)

// JSONBenchmarker interface implements methods that are run for
// benchmarking within BenchmarkSuite
type JSONBenchmarker interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}

// BenchmarkSuite runs a suite of benchmark tests implemented by
// JSONBenchmarker against various datasets
func BenchmarkSuite(b *testing.B, newBenchmarker func() JSONBenchmarker) {

	b.Run("MarshalSmallPayload", func(b *testing.B) {
		bm := newBenchmarker()
		result := SmallPayload{}
		GenerateObjectFromFile(JSON_FILE_SMALL, &result)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bm.Marshal(result)
		}
	})

	b.Run("MarshalLargePayload", func(b *testing.B) {
		bm := newBenchmarker()
		result := LargePayload{}
		GenerateObjectFromFile(JSON_FILE_LARGE, &result)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bm.Marshal(result)
		}
	})

	b.Run("MarshalGeodataPayload", func(b *testing.B) {
		bm := newBenchmarker()
		result := GeoDataPayload{}
		GenerateObjectFromFile(JSON_FILE_GEODATA, &result)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bm.Marshal(result)
		}
	})

	b.Run("UnmarshalSmallPayload", func(b *testing.B) {
		bm := newBenchmarker()
		data := ReadFile(JSON_FILE_SMALL)
		result := &SmallPayload{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bm.Unmarshal(data, result)
		}
	})

	b.Run("UnmarshalLargePayload", func(b *testing.B) {
		bm := newBenchmarker()
		data := ReadFile(JSON_FILE_LARGE)
		result := &LargePayload{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bm.Unmarshal(data, result)
		}
	})

	b.Run("UnmarshalGeodataPayload", func(b *testing.B) {
		bm := newBenchmarker()
		data := ReadFile(JSON_FILE_GEODATA)
		result := &GeoDataPayload{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bm.Unmarshal(data, result)
		}
	})

}
