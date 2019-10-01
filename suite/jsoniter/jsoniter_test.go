// Benchmark suite for jsoniter package
package benchmark_jsoniter

import (
	"github.com/json-iterator/go"
	"github.com/tahia-khan/go-json-benchmark/suite"
	"testing"
)

type Jsoniter struct{}

func (j *Jsoniter) Marshal(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}
func (j *Jsoniter) Unmarshal(b []byte, v interface{}) error {
	return jsoniter.Unmarshal(b, v)
}

func Benchmark(b *testing.B) {
	b.Run("Jsoniter", func(b *testing.B) {
		suite.BenchmarkSuite(b, func() suite.JSONBenchmarker {
			return new(Jsoniter)
		})

	})
}
