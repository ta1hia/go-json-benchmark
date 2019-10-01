// Benchmark suite for ffjson package
package benchmark_ffjson

import (
	"github.com/pquerna/ffjson/ffjson"
	"github.com/tahia-khan/go-json-benchmark/suite"
	"testing"
)

type FfJson struct{}

func (j *FfJson) Marshal(v interface{}) ([]byte, error) {
	return ffjson.MarshalFast(v)
}

func (j *FfJson) Unmarshal(b []byte, v interface{}) error {
	return ffjson.Unmarshal(b, v)
}

func Benchmark(b *testing.B) {
	b.Run("FfJson", func(b *testing.B) {
		suite.BenchmarkSuite(b, func() suite.JSONBenchmarker {
			return new(FfJson)
		})

	})
}
