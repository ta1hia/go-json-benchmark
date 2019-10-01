// Benchmark suite for easyjson package
package benchmark_easyjson

import (
	"github.com/mailru/easyjson"
	"github.com/tahia-khan/go-json-benchmark/suite"

	"testing"
)

type EasyJson struct{}

func (j *EasyJson) Marshal(v interface{}) ([]byte, error) {
	return easyjson.Marshal(v.(easyjson.Marshaler))
}

func (j *EasyJson) Unmarshal(b []byte, v interface{}) error {
	return easyjson.Unmarshal(b, v.(easyjson.Unmarshaler))
}

func Benchmark(b *testing.B) {
	b.Run("EasyJson", func(b *testing.B) {
		suite.BenchmarkSuite(b, func() suite.JSONBenchmarker {
			return new(EasyJson)
		})

	})
}
