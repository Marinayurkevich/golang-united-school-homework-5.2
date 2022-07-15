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
	return Cache{structure: make(map[string]MyCache)}
}

func (f Cache) Get(key string) (string, bool) {
	var value string
	MyCache, ok := f.structure[key]
	if ok {
		expTime := MyCache.expTime
		if expTime {
			value = MyCache.value
			return value, ok
		} else {
			deadline := MyCache.deadline
			timeNow := time.Now()
			expired := deadline.Sub(timeNow).Microseconds()
			if expired > 0 {
				value = MyCache.value
				return value, ok
			} else {
				return value, false
			}
		}
	} else {
		return value, ok
	}
}

func (f Cache) Put(key, value string) {
	deadline := time.Now()
	f.structure[key] = MyCache{
		value:    value,
		expTime:  true,
		deadline: deadline,
	}
}

func (f Cache) Keys() []string {
	var keys []string
	for key, _ := range f.structure {
		_, ok := f.Get(key)
		if ok {
			keys = append(keys, key)
		}
	}
	return keys
}

func (f Cache) PutTill(key, value string, deadline time.Time) {
	f.structure[key] = MyCache{
		value:    value,
		deadline: deadline,
		expTime:  false,
	}
}
