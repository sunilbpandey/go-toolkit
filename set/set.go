package set

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](elems ...T) Set[T] {
	s := make(Set[T])
	for _, elem := range elems {
		s.Add(elem)
	}
	return s
}

func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

func (s Set[T]) Contains(elem T) bool {
	_, exists := s[elem]
	return exists
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Members() []T {
	members := make([]T, 0, len(s))
	for m := range s {
		members = append(members, m)
	}
	return members
}
