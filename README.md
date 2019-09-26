# go-json-benchmark
Go JSON benchmark suite

based on tests in https://github.com/miloyip/nativejson-benchmark

## Libraries

list of libraries this tests

## Benchmarks

### Performance

Benchmark      | Description
---------------|----------------------------------------------------
Marshal        | Parse in-memory JSON into struct.
Unmarshal      | Serialize struct into condensed JSON in memory.
Prettify       | Serialize struct into prettified (with indentation and new lines) JSON in memory.

- run each benchmark for each dataset, if applicable
- generic benchmark helpers in benchmark.go

### Datasets

structs defined in data.go

small.json
large.json
canada.json (geodata)

TODO filesizes + characteristics 


## Adding new libraries for benchmarking

- create LIBNAME_test.go 
- create each benchmark * each dataset - use the base benchmark helpers in benchmark.go if applicable 

TODO
- makefile
- CI/CD
- should i VC the generated files (eg data_easyjson.go, data_ffjson.go) or just include in make instructions?
- should i make the datafixture vars globals rather than reading in each marshal test?
- docstrings and comments
- how to present data
