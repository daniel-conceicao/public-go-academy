package keyValueStore

import (
	"log"
	"sync"

	"github.com/google/uuid"
)

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// KeyValueStore represents the key-value store.
type KeyValueStore struct {
	data map[string]Task
	mu   sync.RWMutex
}

// NewKeyValueStore creates a new instance of KeyValueStore.
func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		data: make(map[string]Task),
	}
}

// Set adds or updates a key-value pair in the store.
func (kv *KeyValueStore) Set(key string, value Task) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	log.Printf("SET params: %s, %v", key, value)
	keyToUse := key
	if keyToUse == "" {
		ok := true
		for ok {
			keyToUse = uuid.NewString()
			_, ok = kv.data[keyToUse]
		}

		value.Id = keyToUse
	}
	kv.data[keyToUse] = value
}

// Get retrieves the value associated with a key from the store.
func (kv *KeyValueStore) Get(key string) (Task, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	val, ok := kv.data[key]
	return val, ok
}

// Deletes the value associated with a key from the store.
func (kv *KeyValueStore) Delete(key string) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	delete(kv.data, key)
}
