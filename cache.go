package cache

import "time"

type Cache struct {
key string
value string
deadline time.Time

}

func NewCache() Cache {
	return Cache{structure: map[string]Cache}
}

func (f Cache) Get(key string) (string, bool) {
if !ok {
return " ", false
}
return Cache.value, true
}

func (f Cache) Put(key, value string) {
}

func (f Cache) Keys() []string {
keys := make ([]string, 0)
for key, value := range NewCache {
keys = append(keys, key)}
return keys
}

func (f Cache) PutTill(key, value string, deadline time.Time) {
}