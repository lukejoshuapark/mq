package functiontypes

type Processor interface {
	Apply(fn func(int) int) error
	Filter(items []string, predicate func(string) bool) []string
	GetHandler() func(string) error
	Transform(data interface{}, mapper func(interface{}) interface{}) interface{}
}
