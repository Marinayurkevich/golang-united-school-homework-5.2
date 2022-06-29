package cache

import "time"

type MyCache struct {
	key      string
	value    string
	deadline time.Time
}

type Cache struct {
	structure map[string]MyCache
}

func NewCache() Cache {
	return Cache{make(map[string]MyCache)}
}

func (f Cache) Get(key string) (string, bool) {
	MyCache, ok := f.structure[key]
	if ok {
		return MyCache.value, true
	}
	return " ", false
}

func (f Cache) Put(key, value string) {

}

func (f Cache) Keys() []string {
	keys := make([]string, 0)
	for key, _ := range f.structure {
		if key >= deadline {
			delete(f.structure, key)
		} else {
			keys = append(keys, key)
		}
	}
	return keys
}

func (f Cache) PutTill(key, value string, deadline time.Time) {

}
