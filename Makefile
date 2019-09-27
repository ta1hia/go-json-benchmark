.PHONY: benchmark build clean

benchmark: 
	go test -bench=. -benchmem

build:
	easyjson -all data.go
	ffjson data.go

readme: build
	go test -bench=. -benchmem > build/results.out
	@python build/rank_results.py

clean: 
	@rm build/results.out
