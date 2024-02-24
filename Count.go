package mq

type Count struct {
	f func(x int) bool
}

func (m Count) ShouldPass(x int) bool {
	return m.f(x)
}

var Never = Count{f: func(x int) bool { return x == 0 }}
var Once = Count{f: func(x int) bool { return x == 1 }}
var AtLeastOnce = Count{f: func(x int) bool { return x >= 1 }}

func Exactly(n int) Count {
	return Count{f: func(x int) bool { return x == n }}
}
