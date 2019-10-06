// Benchmark suite for jingo package. Only tests Marshal
// since jingo provides encoding to json functionality only.
package benchmark_jingo

import (
	"github.com/bet365/jingo"
	"github.com/tahia-khan/go-json-benchmark/suite"
	"testing"
)

// Benchmark runs a benchmarking suite for jingo based on
// the test structure used in suite.BenchmarkSuite.
// Not implementing the JSONBenchmarker interface because
// jingo requires defining an encoder *per* struct.
func Benchmark(b *testing.B) {

	b.Run("Jingo/MarshalSmallPayload", func(b *testing.B) {
		result := suite.SmallPayload{}
		e := jingo.NewStructEncoder(result)
		suite.GenerateObjectFromFile(suite.JSON_FILE_SMALL, &result)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buf := jingo.NewBufferFromPool()
			e.Marshal(&result, buf)
			buf.ReturnToPool()

		}
	})

	// Could not test with LargePayload because the jingo encoder currently
	// does not support marshaling to map[string][string]

	// b.Run("Jingo/MarshalLargePayload", func(b *testing.B) {
	// 	result := suite.LargePayload{}
	// 	e := jingo.NewStructEncoder(result)
	// 	suite.GenerateObjectFromFile(suite.JSON_FILE_SMALL, &result)
	// 	b.ResetTimer()
	// 	for i := 0; i < b.N; i++ {
	// 		buf := jingo.NewBufferFromPool()
	// 		e.Marshal(&result, buf)
	// 		buf.ReturnToPool()
	// 	}
	// })

	b.Run("Jingo/MarshalGeodataPayload", func(b *testing.B) {
		result := suite.GeoDataPayload{}
		e := jingo.NewStructEncoder(result)
		suite.GenerateObjectFromFile(suite.JSON_FILE_SMALL, &result)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buf := jingo.NewBufferFromPool()
			e.Marshal(&result, buf)
			buf.ReturnToPool()
		}
	})

}
