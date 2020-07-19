package cache

// 参考sql的实现
type Cache interface {
	Get(string) (string, error)
	Set(string, string) (bool, error)
	Delete(string) (bool, error)
	SetTimeout(string, int) (bool, error)
}
