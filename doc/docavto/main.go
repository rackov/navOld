package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: gendoc <path>")
		os.Exit(1)
	}

	root := os.Args[1]

	fmt.Println("project/")

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || !info.IsDir() {
			return nil
		}

		base := filepath.Base(path)
		if base == "vendor" || strings.HasPrefix(base, ".") {
			return filepath.SkipDir
		}

		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
		if err != nil || len(pkgs) == 0 {
			return nil
		}

		indent := "├── "
		subIndent := "│   ├── "

		fmt.Println(indent + base)

		for _, pkg := range pkgs {
			for _, f := range pkg.Files {

				for _, decl := range f.Decls {

					// ---------- ФУНКЦИИ И МЕТОДЫ ----------
					if fn, ok := decl.(*ast.FuncDecl); ok && fn.Doc != nil {
						desc := firstLine(fn.Doc.Text())

						if fn.Recv != nil {
							typeName := receiverType(fn)
							fmt.Printf("%s%s.%s — %s\n",
								subIndent,
								typeName,
								fn.Name.Name,
								desc,
							)
						} else {
							fmt.Printf("%s%s — %s\n",
								subIndent,
								fn.Name.Name,
								desc,
							)
						}
					}

					// ---------- ИНТЕРФЕЙСЫ ----------
					gen, ok := decl.(*ast.GenDecl)
					if !ok || gen.Tok != token.TYPE {
						continue
					}

					for _, spec := range gen.Specs {
						ts, ok := spec.(*ast.TypeSpec)
						if !ok {
							continue
						}

						if _, ok := ts.Type.(*ast.InterfaceType); ok && gen.Doc != nil {
							desc := firstLine(gen.Doc.Text())
							fmt.Printf("%s%s — %s\n",
								subIndent,
								ts.Name.Name,
								desc,
							)
						}
					}
				}
			}
		}

		return nil
	})
}

func firstLine(text string) string {
	return strings.TrimSpace(strings.Split(text, "\n")[0])
}

func receiverType(fn *ast.FuncDecl) string {
	recv := fn.Recv.List[0].Type
	switch t := recv.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		if id, ok := t.X.(*ast.Ident); ok {
			return id.Name
		}
	}
	return "unknown"
}
