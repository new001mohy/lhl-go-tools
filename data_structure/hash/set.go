package hash

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable](items ...T) *Set[T] {
	set := &Set[T]{m: make(map[T]struct{})}
	for _, item := range items {
		set.Add(item)
	}
	return set
}

func NewSetWithCapacity[T comparable](capacity int, items ...T) *Set[T] {
	set := &Set[T]{m: make(map[T]struct{}, capacity)}
	for _, item := range items {
		set.Add(item)
	}
	return set
}

func (s *Set[T]) Add(item T) {
	s.m[item] = struct{}{}
}
func (s *Set[T]) Remove(item T) {
	delete(s.m, item)
}
func (s *Set[T]) Contains(item T) bool {
	_, ok := s.m[item]
	return ok
}
func (s *Set[T]) Size() int {
	return len(s.m)
}
func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}
func (s *Set[T]) Values() []T {
	values := make([]T, 0, s.Size())
	for item := range s.m {
		values = append(values, item)
	}
	return values
}
func (s *Set[T]) Equals(otherSet *Set[T]) bool {
	if otherSet == nil || len(s.m) != len(otherSet.m) {
		return false
	}
	for item := range s.m {
		if !otherSet.Contains(item) {
			return false
		}
	}
	return true
}
