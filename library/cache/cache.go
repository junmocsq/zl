package cache

// 参考sql的实现
type Cache interface {
	Get(string) (string, error)
	Set(string, string) (bool, error)
	SetEx(string, string, int) (bool, error)
	Delete(string) (bool, error)
	SetTimeout(string, int) (bool, error)
}

var cache Cache

func NewCache() Cache {
	if cache == nil {
		panic("请注册缓存所需组件")
	}
	return cache
}

func RegisterCache(c Cache) {
	cache = c
}
