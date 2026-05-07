package erroronly

type Store interface {
	Save(key string, value string) error
	Delete(key string) error
	Connect() error
}
