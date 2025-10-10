package noparameters

type StatusProvider interface {
	GetStatus() string
	IsReady() (bool, error)
	Refresh()
}
