// Benchmark suite for encoding/json package (referred to as 'standard')
package benchmark_standard

import (
	"encoding/json"
	"github.com/tahia-khan/go-json-benchmark/suite"
	"testing"
)

type StandardJson struct{}

func (j *StandardJson) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func (j *StandardJson) Unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

func Benchmark(b *testing.B) {
	b.Run("StandardJson", func(b *testing.B) {
		suite.BenchmarkSuite(b, func() suite.JSONBenchmarker {
			return new(StandardJson)
		})

	})
}
