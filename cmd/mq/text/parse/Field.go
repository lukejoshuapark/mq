package parse

import (
	"go/ast"
	"go/types"
)

type Field struct {
	Name string
	Type string
}

func ParseFields(astField *ast.Field) []Field {
	t := types.ExprString(astField.Type)
	if len(astField.Names) < 1 {
		return []Field{
			{
				Name: "",
				Type: t,
			},
		}
	}

	var fields []Field
	for _, name := range astField.Names {
		fields = append(fields, Field{
			Name: name.Name,
			Type: t,
		})
	}

	return fields
}
