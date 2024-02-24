package parse

import "go/ast"

type Method struct {
	Name    string
	Inputs  []Field
	Outputs []Field
}

func ParseMethod(method *ast.Field) Method {
	return Method{
		Name:    getMethodName(method),
		Inputs:  getMethodInputs(method),
		Outputs: getMethodOutputs(method),
	}
}

func getMethodName(method *ast.Field) string {
	return method.Names[0].Name
}

func getMethodInputs(method *ast.Field) []Field {
	var inputs []Field

	for _, field := range method.Type.(*ast.FuncType).Params.List {
		inputs = append(inputs, ParseFields(field)...)
	}

	return inputs
}

func getMethodOutputs(method *ast.Field) []Field {
	if method.Type.(*ast.FuncType).Results == nil {
		return nil
	}

	var outputs []Field

	for _, field := range method.Type.(*ast.FuncType).Results.List {
		outputs = append(outputs, ParseFields(field)...)
	}

	return outputs
}
