package importgroup

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "importgroup",
	Doc:  Doc,
	Run:  run,
}

const Doc = "importgroup finds multiple import statements which are without grouping"

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		var count int

		for _, decl := range f.Decls {
			gendecl, ok := decl.(*ast.GenDecl)
			if !ok || gendecl.Tok != token.IMPORT {
				continue
			}

			count++

			if count < 2 {
				continue
			}

			for _, spec := range gendecl.Specs {
				if im, ok := spec.(*ast.ImportSpec); ok {
					pass.Reportf(im.Pos(), "the import statement is not grouped")
				}
			}
		}
	}
	return nil, nil
}
