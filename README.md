# go-json-benchmark
Benchmark suite for JSON parsing and serialization. 

## Libraries

Library | Notes
--------|-------------------
[encoding/json (aka 'std')](http://golang.org/pkg/encoding/json/) | 
[easyjson](http://github.com/mailru/easyjson/) | 
[ffjson](http://github.com/pquerna/ffjson/) | 
[jsoniter](http://github.com/json-iterator/go/) | 

## Benchmarks

### Performance

Benchmark      | Description
---------------|----------------------------------------------------
Marshal        | Encode JSON data from a struct
Unmarshal      | Decode JSON data into to a struct
Prettify       | Encode JSON data from a struct into prettified JSON (with indentation and new lines)


Each benchmark should be tested against each dataset. Generic benchmark helpers are defined in benchmark.go

### Datasets

JSON file   | Size | Description
------------|------|-----------------------
`small.json` [source](https://github.com/tahia-khan/go-json-benchmark/blob/master/data/small.json) | 1KB | 
`large.json` [source](https://github.com/miloyip/nativejson-benchmark/blob/master/data/citm_catalog.json) | 1687KB | A big benchmark file with indentation used in several Java JSON parser benchmarks.
`canada.json` [source](https://github.com/miloyip/nativejson-benchmark/blob/master/data/canada.json) | 2199KB | Contour of Canada border in [GeoJSON](http://geojson.org/) format. Contains a lot of real numbers.

These datasets were pulled from [nativejson-benchmark](https://github.com/miloyip/nativejson-benchmark)

## Adding new libraries for benchmarking

- create LIBNAME_test.go 
- create each benchmark * each dataset - use the base benchmark helpers in benchmark.go if applicable 

TODO
- makefile
- CI/CD
- data presentation
