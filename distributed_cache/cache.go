package distributed_cache

import (
	"context"
	"fmt"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key, value string, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string)
}

type MockCache struct {
	data map[string]mockEntry
}

type mockEntry struct {
	value     string
	expiresAt time.Time
}

func NewMockCache() *MockCache {
	return &MockCache{data: make(map[string]mockEntry)}
}

func (m *MockCache) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	m.data[key] = mockEntry{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
	return nil
}

func (m *MockCache) Get(ctx context.Context, key string) (string, error) {
	entry, found := m.data[key]
	if !found || time.Now().After(entry.expiresAt) {
		return "", fmt.Errorf("key not found or expired")
	}
	return entry.value, nil
}

func (m *MockCache) Delete(ctx context.Context, key string) {
	delete(m.data, key)
}
