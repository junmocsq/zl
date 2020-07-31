package cache

// 参考sql的实现
type Cache interface {
	Get(string) (string, error)
	Set(string, string) (bool, error)
	SetEx(string, string, int) (bool, error)
	Delete(string) (bool, error)
	SetTimeout(string, int) (bool, error)
}

func NewCache() *redis {
	return &redis{}
}
