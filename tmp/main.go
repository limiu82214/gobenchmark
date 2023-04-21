package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	// 解析 hello.go 的 AST
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "./tmp/main/hello.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	ast.Print(fset, file)
}
