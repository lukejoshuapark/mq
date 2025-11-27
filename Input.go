package mq

import (
	"fmt"
	"reflect"
)

type Input[T any] struct {
	v           T
	f           func(v T) bool
	description string
}

func (i Input[T]) Compare(v T) bool {
	if i.f == nil {
		return reflect.DeepEqual(i.v, v)
	}

	return i.f(v)
}

func (i Input[T]) String() string {
	if i.description != "" {
		return i.description
	}

	if i.f == nil {
		return fmt.Sprintf("%+v", i.v)
	}

	return "(custom matcher)"
}

func (i Input[T]) WithDescription(description string) Input[T] {
	i.description = description
	return i
}

func IsExactly[T any](v T) Input[T] {
	return Input[T]{
		v: v,
	}
}

func IsAny[T any]() Input[T] {
	return Input[T]{
		f: func(v T) bool {
			return true
		},
	}
}

func Is[T any](f func(v T) bool) Input[T] {
	return Input[T]{
		f: f,
	}
}
