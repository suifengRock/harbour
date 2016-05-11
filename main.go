package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src/apidoc/init.go", nil, 0)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.FuncDecl:
			fmt.Println(t.Name.Name)
			for _, stmt := range t.Body.List {
				call := stmt.(*ast.ExprStmt).X.(*ast.CallExpr)
				fun := call.Fun.(*ast.SelectorExpr)
				fmt.Println(fun.X.(*ast.Ident).Name, fun.Sel.Name)
				for _, arg := range call.Args {

				}
			}
			return false
		}
		return true
	})

	// Print the AST.
	ast.Print(fset, f)
}
