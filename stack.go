package main

// Stack represents a generic stack data structure.
type Stack[T any] struct {
	items []T
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Peek returns the item at the top of the stack.
// If the stack is empty, it returns the zero value of type T.
func (s *Stack[T]) Peek() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	return s.items[len(s.items)-1]
}

// Pop removes and returns the item at the top of the stack.
// If the stack is empty, it returns the zero value of type T.
func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Count returns the number of items in the stack.
func (s *Stack[T]) Count() int {
	return len(s.items)
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
