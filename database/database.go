package database

import (
	"fmt"
	"sync/atomic"
)

type Template struct {
	Key  string
	Name string
}

var counter int64

type TemplateRepository interface {
	GetTemplateByID(key string) (*Template, error)
}

type MockTemplateRepository struct{}

func NewMockTemplateRepository() TemplateRepository {
	return &MockTemplateRepository{}
}

func (m *MockTemplateRepository) GetTemplateByID(key string) (*Template, error) {
	atomic.AddInt64(&counter, 1)
	fmt.Println("database call hit: ", atomic.LoadInt64(&counter))
	mockData := map[string]*Template{
		"key1": {Key: "key1", Name: "test-1"},
		"key2": {Key: "key1", Name: "test-2"},
	}
	if template, exists := mockData[key]; exists {
		return template, nil
	}
	return nil, fmt.Errorf("template not found")
}
