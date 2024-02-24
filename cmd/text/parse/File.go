package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type File struct {
	PackageName string
	Imports     []string // TODO: This.
	Interfaces  []Interface
}

func ParseFile(fileName string) (File, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fileName, nil, parser.AllErrors)
	if err != nil {
		return File{}, err
	}

	var interfaces []Interface

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			interfaceType, ok := typeSpec.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}

			interfaces = append(interfaces, ParseInterface(typeSpec.Name.Name, interfaceType))
		}
	}

	return File{
		PackageName: node.Name.Name,
		Imports:     getImports(node),
		Interfaces:  interfaces,
	}, nil
}

func getImports(node *ast.File) []string {
	var imports []string

	for _, imp := range node.Imports {
		imports = append(imports, strings.Trim(imp.Path.Value, "\""))
	}

	return imports
}
