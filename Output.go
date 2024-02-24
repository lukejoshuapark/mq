package mq

type Output[T any] struct {
	v T
	f func() T
	p *T
}

func (o Output[T]) Value() T {
	if o.f != nil {
		return o.f()
	}

	if o.p != nil {
		return *o.p
	}

	return o.v
}

func Returns[T any](p *T) Output[T] {
	return Output[T]{
		p: p,
	}
}

func ReturnsExactly[T any](v T) Output[T] {
	return Output[T]{
		v: v,
	}
}

func ReturnsDefault[T any]() Output[T] {
	return Output[T]{}
}

func ReturnsOutputFrom[T any](f func() T) Output[T] {
	return Output[T]{
		f: f,
	}
}
