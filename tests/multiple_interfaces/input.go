package multiinterface

type Reader interface {
	Read(key string) (string, error)
	Exists(key string) bool
}

type Writer interface {
	Write(key string, value string) error
	Delete(key string) error
}
