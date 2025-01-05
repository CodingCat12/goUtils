package main

type Selection[T comparable] map[T]struct{}

func (s Selection[T]) Toggle(key T) {
	if _, ok := s[key]; ok {
		s.Remove(key)
	} else {
		s.Add(key)
	}
}

func (s Selection[T]) Contains(key T) bool {
	_, ok := s[key]
	return ok
}

func (s Selection[T]) Add(key T) {
	s[key] = struct{}{}
}

func (s Selection[T]) Remove(key T) {
	delete(s, key)
}

func (s Selection[T]) Clear() {
	for key := range s {
		delete(s, key)
	}
}
