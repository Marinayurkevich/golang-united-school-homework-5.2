package cache

import "time"

type MyCache struct {
	value    string
	deadline time.Time
	expTime  bool
}

type Cache struct {
	structure map[string]MyCache
}

func NewCache() Cache {
	return Cache{make(map[string]MyCache)}
}

func (f Cache) Get(key string) (string, bool) {
	MyCache, ok := f.structure[key]
	expTime := MyCache.expTime
	if ok || expTime {
		value := MyCache.value
		return value, true
	}
	deadline := MyCache.deadline
	currentTime := time.Now()
	evaluate := deadline.Sub(currentTime)
	if evaluate <= 0 {
		return " ", false
	} else {
		value := MyCache.value
		return value, true
	}

}

func (f Cache) Put(key, value string) {
	f.structure[key] = MyCache{
		expTime: true,
	}
}

func (f Cache) Keys() []string {
	keys := make([]string, 0)
	for key, _ := range f.structure {
		keys = append(keys, key)
	}
	return keys
}

func (f Cache) PutTill(key, value string, deadline time.Time) {
	f.structure[key] = MyCache{}
}
