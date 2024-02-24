package parse

import (
	"go/ast"
)

type Interface struct {
	Name    string
	Methods []Method
}

func ParseInterface(name string, interfaceType *ast.InterfaceType) Interface {
	var methods []Method

	for _, method := range interfaceType.Methods.List {
		methods = append(methods, ParseMethod(method))
	}

	return Interface{
		Name:    name,
		Methods: methods,
	}
}
