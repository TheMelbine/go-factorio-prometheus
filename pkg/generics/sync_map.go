package generics

import (
	"iter"
	"sync"
)

type SyncMap[K any, V any] struct {
	data sync.Map
}

func NewSyncMap[K, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		data: sync.Map{},
	}
}

func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.data.Load(key)
	if !ok {
		return
	}
	value = v.(V)
	return
}

func (m *SyncMap[K, V]) Store(key K, value V) {
	m.data.Store(key, value)
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.data.Delete(key)
}

func (m *SyncMap[K, V]) Items() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		m.data.Range(func(key, value interface{}) bool {
			return yield(key.(K), value.(V))
		})
	}
}