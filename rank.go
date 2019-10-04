package main

import (
	"github.com/tahia-khan/go-json-benchmark/internal/rank"
)

func main() {
	rank.GenerateREADME("results.out", "README.md.template", "README.md")
}
