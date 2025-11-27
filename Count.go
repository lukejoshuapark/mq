package mq

import "fmt"

type Count struct {
	f           func(x int) bool
	description string
}

func (m Count) ShouldPass(x int) bool {
	return m.f(x)
}

func (m Count) String() string {
	if m.description != "" {
		return m.description
	}

	return "(custom count)"
}

func (m Count) WithDescription(description string) Count {
	m.description = description
	return m
}

var Never = Count{f: func(x int) bool { return x == 0 }, description: "never"}
var Once = Count{f: func(x int) bool { return x == 1 }, description: "once"}
var AtLeastOnce = Count{f: func(x int) bool { return x >= 1 }, description: "at least once"}

func Exactly(n int) Count {
	return Count{f: func(x int) bool { return x == n }, description: fmt.Sprintf("exactly %d times", n)}
}
