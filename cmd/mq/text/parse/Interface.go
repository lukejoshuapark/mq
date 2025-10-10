package parse

import (
	"fmt"
	"go/ast"
)

type Interface struct {
	Name    string
	Methods []Method
}

func ParseInterface(name string, interfaceType *ast.InterfaceType) (Interface, error) {
	var methods []Method

	for _, method := range interfaceType.Methods.List {
		m, err := ParseMethod(method)
		if err != nil {
			return Interface{}, fmt.Errorf("failed to parse method in interface %s: %w", name, err)
		}
		methods = append(methods, m)
	}

	return Interface{
		Name:    name,
		Methods: methods,
	}, nil
}
