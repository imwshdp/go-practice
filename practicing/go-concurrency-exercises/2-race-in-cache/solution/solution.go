package solution

import (
	"container/list"
	"sync"
	"testing"
)

const CacheSize = 100

type KeyStoreCacheLoader interface {
	Load(string) string
}

type page struct {
	Key   string
	Value string
}

type KeyStoreCache struct {
	mu    sync.Mutex
	cache map[string]*list.Element
	pages list.List
	load  func(string) string
}

func New(load KeyStoreCacheLoader) *KeyStoreCache {
	return &KeyStoreCache{
		load:  load.Load,
		cache: make(map[string]*list.Element),
	}
}

func (k *KeyStoreCache) Get(key string) string {
	k.mu.Lock()
	defer k.mu.Unlock()

	if e, ok := k.cache[key]; ok {
		k.pages.MoveToFront(e)
		value := e.Value.(page).Value
		return value
	}

	p := page{key, k.load(key)}

	if len(k.cache) >= CacheSize {
		end := k.pages.Back()
		delete(k.cache, end.Value.(page).Key)
		k.pages.Remove(end)
	}
	k.pages.PushFront(p)
	k.cache[key] = k.pages.Front()

	return p.Value
}

type Loader struct {
	DB *MockDB
}

func (l *Loader) Load(key string) string {
	val, err := l.DB.Get(key)
	if err != nil {
		panic(err)
	}

	return val
}

func run(t *testing.T) (*KeyStoreCache, *MockDB) {
	loader := Loader{
		DB: GetMockDB(),
	}
	cache := New(&loader)

	RunMockServer(cache, t)

	return cache, loader.DB
}

func main() {
	run(nil)
}
