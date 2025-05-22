package main

import (
	"sync"
)

// Store is a thread-safe in-memory key-value store.
// It uses a map to store data and RWMutex to handle concurrent access.
type Store struct {
	sync.RWMutex      // Embed RWMutex for read/write locking
	data map[string]string // Internal map to store key-value pairs
}

// NewStore creates and returns a new instance of Store.
// The internal map is initialized.
func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

// Set stores a key-value pair in the store.
// It acquires a write lock to ensure data consistency during concurrent access.
func (s *Store) Set(key string, value string) {
	s.Lock()         // Acquire write lock
	defer s.Unlock() // Ensure lock is released
	s.data[key] = value
}

// Get retrieves the value for a given key from the store.
// It returns the value and true if the key exists, otherwise an empty string and false.
// It acquires a read lock, allowing multiple concurrent reads.
func (s *Store) Get(key string) (string, bool) {
	s.RLock()         // Acquire read lock
	defer s.RUnlock() // Ensure lock is released
	val, ok := s.data[key]
	return val, ok
}

// Delete removes a key and its associated value from the store.
// It acquires a write lock to ensure data consistency during concurrent access.
func (s *Store) Delete(key string) {
	s.Lock()         // Acquire write lock
	defer s.Unlock() // Ensure lock is released
	delete(s.data, key)
}
