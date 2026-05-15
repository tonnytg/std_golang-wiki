package cache

type Cacheable interface {
	entities.Category | entites.Post
}

type Cache[T Cacheable] struct {
	data map[string]T
}

func (c *Cache[T]) Set(key string, value T) {
	c.data[key] = value
}

func (c *Cache[T]) Get(key string) (v V) {
	if v, ok := c.data[key]; ok {
		return v
	}
	return
}

func New[T Cacheable]() *cache[T] {
	c := cache[T]{}
	c.data = make(map[string]T)

	return &c
}
