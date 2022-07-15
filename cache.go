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
	var value string
	MyCache, ok := f.structure[key]
	expTime := MyCache.expTime
	if ok {
		expTime = MyCache.expTime
		if expTime {
			value := MyCache.value
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
	f.structure[key] = MyCache{
		expTime: true,
		value:   value,
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
	f.structure[key] = MyCache{
		expTime:  true,
		value:    value,
		deadline: deadline,
	}
}
