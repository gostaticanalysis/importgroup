package importgroup_test

import (
	"testing"

	"github.com/gostaticanalysis/importgroup"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, importgroup.Analyzer, "a")
}