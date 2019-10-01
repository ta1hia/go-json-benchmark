.PHONY: benchmark build clean

benchmark: 
	go test -bench=. ./... -benchmem

build:
	easyjson -all suite/data.go
	ffjson suite/data.go
