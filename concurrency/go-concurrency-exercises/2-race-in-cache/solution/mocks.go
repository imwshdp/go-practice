package solution

import "testing"

type MockDB struct {
	Calls int32
}

func (db *MockDB) Get(key string) (string, error) {
	return "", nil
}

func GetMockDB() *MockDB {
	return nil
}

func RunMockServer(cache *KeyStoreCache, t *testing.T) {}
