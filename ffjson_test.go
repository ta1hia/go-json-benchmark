package json_benchmark

import (
	"github.com/pquerna/ffjson/ffjson"
	"testing"
)

func BenchmarkMarshalSmallPayloadByFFJson(b *testing.B) {
	result := SmallPayload{}
	GenerateObjectFromFile(JSON_FILE_SMALL, &result)
	benchmarkMarshalGenericData(result, ffjson.MarshalFast, b)
}

func BenchmarkMarshalLargePayloadByFFJson(b *testing.B) {
	result := LargePayload{}
	GenerateObjectFromFile(JSON_FILE_LARGE, &result)
	benchmarkMarshalGenericData(result, ffjson.MarshalFast, b)
}

func BenchmarkMarshalGeoDataPayloadByFFJson(b *testing.B) {
	result := GeoDataPayload{}
	GenerateObjectFromFile(JSON_FILE_GEODATA, &result)
	benchmarkMarshalGenericData(result, ffjson.MarshalFast, b)
}

func BenchmarkUnmarshalSmallPayloadByFFJson(b *testing.B) {
	data := ReadFile(JSON_FILE_SMALL)
	result := SmallPayload{}
	benchmarkUnmarshalGenericData(data, &result, ffjson.UnmarshalFast, b)
}

func BenchmarkUnmarshalLargePayloadByFFJson(b *testing.B) {
	data := ReadFile(JSON_FILE_LARGE)
	result := LargePayload{}
	benchmarkUnmarshalGenericData(data, &result, ffjson.UnmarshalFast, b)
}

func BenchmarkUnmarshalGeoDataPayloadByFFJson(b *testing.B) {
	data := ReadFile(JSON_FILE_GEODATA)
	result := GeoDataPayload{}
	benchmarkUnmarshalGenericData(data, &result, ffjson.UnmarshalFast, b)
}
