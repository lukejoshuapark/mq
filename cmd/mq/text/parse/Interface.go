package parse

import (
	"fmt"
	"go/ast"
	"go/types"
)

type Interface struct {
	Name       string
	TypeParams string // e.g., "[T any]" or "[K comparable, V any]"
	Methods    []Method
}

func ParseInterface(name string, typeParams *ast.FieldList, interfaceType *ast.InterfaceType) (Interface, error) {
	var methods []Method

	for _, method := range interfaceType.Methods.List {
		m, err := ParseMethod(method)
		if err != nil {
			return Interface{}, fmt.Errorf("failed to parse method in interface %s: %w", name, err)
		}
		methods = append(methods, m)
	}

	typeParamsStr := ""
	if typeParams != nil && len(typeParams.List) > 0 {
		typeParamsStr = formatTypeParams(typeParams)
	}

	return Interface{
		Name:       name,
		TypeParams: typeParamsStr,
		Methods:    methods,
	}, nil
}

func formatTypeParams(params *ast.FieldList) string {
	if params == nil || len(params.List) == 0 {
		return ""
	}

	result := "["
	for i, param := range params.List {
		if i > 0 {
			result += ", "
		}

		for j, name := range param.Names {
			if j > 0 {
				result += ", "
			}
			result += name.Name
		}

		result += " " + types.ExprString(param.Type)
	}
	result += "]"

	return result
}
