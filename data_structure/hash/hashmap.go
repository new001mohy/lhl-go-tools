package hash

type Map[k comparable, v any] struct {
	data map[k]v
}

func NewMap[K comparable, v any]() *Map[K, v] {
	return &Map[K, v]{
		data: make(map[K]v, 8),
	}
}

func NewMapWithCapacity[K comparable, v any](capacity int) *Map[K, v] {
	return &Map[K, v]{
		data: make(map[K]v, capacity),
	}
}

func (m *Map[K, v]) Put(key K, value any) {
	m.data[key] = value
}

func (m *Map[K, v]) Get(key K) (v any, ok bool) {
	v, ok = m.data[key]
	return v, ok

}
func (m *Map[K, v]) Remove(key K) {
	delete(m.data, key)
}
func (m *Map[K, v]) Len() int {
	return len(m.data)
}
func (m *Map[K, v]) Keys() []K {
	keys := make([]K, 0, m.Len())
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}
func (m *Map[K, v]) Values() []any {
	values := make([]any, 0, m.Len())
	for _, v := range m.data {
		values = append(values, v)
	}
	return values
}
func (m *Map[K, v]) Clear() {
	m.data = make(map[K]v)
}
func (m *Map[K, v]) ContainsKey(key K) bool {
	_, ok := m.data[key]
	return ok
}
