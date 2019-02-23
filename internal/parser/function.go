package parser

import (
	"genopi/internal/common"
	"go/ast"
)

func readMethod(f *ast.File, x1 *ast.FuncDecl) (common.Method, bool) {
	if x1.Doc != nil {
		comments := make([]string, 0)
		for _, d := range x1.Doc.List {
			comments = append(comments, d.Text[3:])
		}
		return common.Method{
			Package:  f.Name.Name,
			Name:     x1.Name.Name,
			Comments: comments,
		}, true
	}

	return common.Method{}, false
}
