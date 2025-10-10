package parse

import (
	"fmt"
	"go/ast"
)

type Method struct {
	Name    string
	Inputs  []Field
	Outputs []Field
}

func ParseMethod(method *ast.Field) (Method, error) {
	name, err := getMethodName(method)
	if err != nil {
		return Method{}, err
	}

	inputs, err := getMethodInputs(method)
	if err != nil {
		return Method{}, err
	}

	outputs, err := getMethodOutputs(method)
	if err != nil {
		return Method{}, err
	}

	return Method{
		Name:    name,
		Inputs:  inputs,
		Outputs: outputs,
	}, nil
}

func getMethodName(method *ast.Field) (string, error) {
	if len(method.Names) == 0 {
		return "", fmt.Errorf("method has no name (possibly an embedded interface)")
	}
	return method.Names[0].Name, nil
}

func getMethodInputs(method *ast.Field) ([]Field, error) {
	if method.Type == nil {
		return nil, fmt.Errorf("method type is nil")
	}

	funcType, ok := method.Type.(*ast.FuncType)
	if !ok {
		return nil, fmt.Errorf("method type is not a FuncType")
	}

	if funcType.Params == nil {
		return nil, nil
	}

	var inputs []Field
	for _, field := range funcType.Params.List {
		inputs = append(inputs, ParseFields(field)...)
	}

	return inputs, nil
}

func getMethodOutputs(method *ast.Field) ([]Field, error) {
	if method.Type == nil {
		return nil, fmt.Errorf("method type is nil")
	}

	funcType, ok := method.Type.(*ast.FuncType)
	if !ok {
		return nil, fmt.Errorf("method type is not a FuncType")
	}

	if funcType.Results == nil {
		return nil, nil
	}

	var outputs []Field
	for _, field := range funcType.Results.List {
		outputs = append(outputs, ParseFields(field)...)
	}

	return outputs, nil
}
