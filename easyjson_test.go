package json_benchmark

import (
	"testing"
)

func BenchmarkMarshalSmallPayloadByEasyJson(b *testing.B) {
	result := SmallPayload{}
	GenerateObjectFromFile(JSON_FILE_SMALL, &result)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.MarshalJSON()
	}
}

func BenchmarkMarshalLargePayloadByEasyJson(b *testing.B) {
	result := LargePayload{}
	GenerateObjectFromFile(JSON_FILE_LARGE, &result)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.MarshalJSON()
	}
}

func BenchmarkMarshalGeoDataPayloadByEasyJson(b *testing.B) {
	result := GeoDataPayload{}
	GenerateObjectFromFile(JSON_FILE_GEODATA, &result)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.MarshalJSON()
	}
}

func BenchmarkUnmarshalSmallPayloadByEasyJson(b *testing.B) {
	data := ReadFile(JSON_FILE_SMALL)
	result := &SmallPayload{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.UnmarshalJSON(data)
	}
}

func BenchmarkUnmarshalLargePayloadByEasyJson(b *testing.B) {
	data := ReadFile(JSON_FILE_LARGE)
	result := &LargePayload{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.UnmarshalJSON(data)
	}
}

func BenchmarkUnmarshalGeoDataPayloadByEasyJson(b *testing.B) {
	data := ReadFile(JSON_FILE_GEODATA)
	result := &GeoDataPayload{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.UnmarshalJSON(data)
	}
}
