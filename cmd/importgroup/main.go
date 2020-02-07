package main

import (
	"github.com/gostaticanalysis/importgroup"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(importgroup.Analyzer) }
