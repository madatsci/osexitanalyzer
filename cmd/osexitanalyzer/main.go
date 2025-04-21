package main

import (
	"github.com/madatsci/osexitanalyzer"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		osexitanalyzer.Analyzer,
	)
}
