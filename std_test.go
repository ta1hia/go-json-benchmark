package json_benchmark

import (
	"encoding/json"
	"testing"
)

func BenchmarkMarshalSmallPayloadByStandardJson(b *testing.B) {
	result := SmallPayload{}
	GenerateObjectFromFile(JSON_FILE_SMALL, &result)
	benchmarkMarshalGenericData(result, json.Marshal, b)
}

func BenchmarkMarshalLargePayloadByStandardJson(b *testing.B) {
	result := LargePayload{}
	GenerateObjectFromFile(JSON_FILE_LARGE, &result)
	benchmarkMarshalGenericData(result, json.Marshal, b)
}

func BenchmarkMarshalGeoDataPayloadByStandardJson(b *testing.B) {
	result := GeoDataPayload{}
	GenerateObjectFromFile(JSON_FILE_GEODATA, &result)
	benchmarkMarshalGenericData(result, json.Marshal, b)
}

func BenchmarkUnmarshalSmallPayloadByStandardJson(b *testing.B) {
	data := ReadFile(JSON_FILE_SMALL)
	result := SmallPayload{}
	benchmarkUnmarshalGenericData(data, &result, json.Unmarshal, b)
}

func BenchmarkUnmarshalLargePayloadByStandardJson(b *testing.B) {
	data := ReadFile(JSON_FILE_LARGE)
	result := LargePayload{}
	benchmarkUnmarshalGenericData(data, &result, json.Unmarshal, b)
}

func BenchmarkUnmarshalGeoDataPayloadByStandardJson(b *testing.B) {
	data := ReadFile(JSON_FILE_GEODATA)
	result := GeoDataPayload{}
	benchmarkUnmarshalGenericData(data, &result, json.Unmarshal, b)
}
