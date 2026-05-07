package maptypes

type Registry interface {
	Set(key string, tags map[string]string) error
	Get(key string) (map[string]string, error)
	Merge(base map[string]string, overlay map[string]string) map[string]string
}
