[![Build Status](https://travis-ci.org/tahia-khan/go-json-benchmark.svg?branch=master)](https://travis-ci.org/tahia-khan/go-json-benchmark)

# go-json-benchmark
Benchmark suite for JSON parsing and serialization. 

## Libraries

Library | Notes
--------|-------------------
[encoding/json (aka 'std')](http://golang.org/pkg/encoding/json/) | Testing Marshal/Unmarshal rather than Decode/Encode since this benchmark suite assumes JSON exists in memory and not incoming from a stream
[easyjson](http://github.com/mailru/easyjson/) | Generates parser methods structs defined in `<data_file>.go`. Testing MarshalJSON/UnmarshalJSON.
[ffjson](http://github.com/pquerna/ffjson/) | Generates parser methods structs defined in `<data_file>.go`. Testing MarshalFast/UnmarshalFast
[jsoniter](http://github.com/json-iterator/go/) | Testing Marshal/Unmarshal.
[jingo](http://github.com/bet365/jingo/) | Testing Marshal (no Unmarshal). Requires initializing an encoder per struct/slice that should be marshaled. Currently does not marshal maps.

## Benchmarks

### Performance

Benchmark      | Description
---------------|----------------------------------------------------
Marshal        | Encode JSON data from a Go type or struct
Unmarshal      | Decode JSON data from memory into to a Go type or struct

Each benchmark should be tested against each dataset. Generic benchmark helpers are defined in `benchmark.go`.

### Datasets

JSON file   | Size | Description
------------|------|-----------------------
`small.json` [source](https://github.com/tahia-khan/go-json-benchmark/blob/master/data/small.json) | 1KB | 
`large.json` [source](https://github.com/miloyip/nativejson-benchmark/blob/master/data/citm_catalog.json) | 1687KB | A big benchmark file with indentation used in several Java JSON parser benchmarks.
`canada.json` [source](https://github.com/miloyip/nativejson-benchmark/blob/master/data/canada.json) | 2199KB | Contour of Canada border in [GeoJSON](http://geojson.org/) format. Contains a lot of real numbers.

These datasets were pulled from [nativejson-benchmark](https://github.com/miloyip/nativejson-benchmark). Data structs are defined in `data.go`.

## Results

Results of Marshal benchmarking:

Rank | Library | Performance
-----|---------|---------
1 | jingo | 326 ns/op
2 | ffjson | 327 ns/op
3 | easyjson | 12379009 ns/op
4 | standardjson | 26374413 ns/op
5 | jsoniter | 26768552 ns/op

Results of Unmarshal benchmarking:

Rank | Library | Performance
-----|---------|---------
1 | easyjson | 20851164 ns/op
2 | ffjson | 21516266 ns/op
3 | jsoniter | 27719141 ns/op
4 | standardjson | 39032163 ns/op

The average across all datasets is calculated for each performance metric.

## Adding new libraries for benchmarking

- create LIBNAME_test.go 
- create each benchmark * each dataset - use the base benchmark helpers in benchmark.go if applicable 

## TODO
- makefile
- CI/CD
- data presentation
- benchmark for libs that support field querying rather than converting to/from predefined struct
