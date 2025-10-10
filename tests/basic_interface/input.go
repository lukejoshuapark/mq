package basic

type SimpleService interface {
	DoSomething(input string) (string, error)
}
