package json_benchmark

import (
	"github.com/json-iterator/go"
	"testing"
)

func BenchmarkMarshalSmallPayloadByJsoniter(b *testing.B) {
	result := SmallPayload{}
	GenerateObjectFromFile(JSON_FILE_SMALL, &result)
	benchmarkMarshalGenericData(result, jsoniter.Marshal, b)
}

func BenchmarkMarshalLargePayloadByJsoniter(b *testing.B) {
	result := LargePayload{}
	GenerateObjectFromFile(JSON_FILE_LARGE, &result)
	benchmarkMarshalGenericData(result, jsoniter.Marshal, b)
}

func BenchmarkMarshalGeoDataPayloadByJsoniter(b *testing.B) {
	result := GeoDataPayload{}
	GenerateObjectFromFile(JSON_FILE_GEODATA, &result)
	benchmarkMarshalGenericData(result, jsoniter.Marshal, b)
}

func BenchmarkUnmarshalSmallPayloadByJsoniter(b *testing.B) {
	data := ReadFile(JSON_FILE_SMALL)
	result := SmallPayload{}
	benchmarkUnmarshalGenericData(data, &result, jsoniter.Unmarshal, b)
}

func BenchmarkUnmarshalLargePayloadByJsoniter(b *testing.B) {
	data := ReadFile(JSON_FILE_LARGE)
	result := LargePayload{}
	benchmarkUnmarshalGenericData(data, &result, jsoniter.Unmarshal, b)
}

func BenchmarkUnmarshalGeoDataPayloadByJsoniter(b *testing.B) {
	data := ReadFile(JSON_FILE_GEODATA)
	result := GeoDataPayload{}
	benchmarkUnmarshalGenericData(data, &result, jsoniter.Unmarshal, b)
}
