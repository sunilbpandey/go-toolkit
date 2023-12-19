package set

import (
	"slices"
	"sort"
	"testing"
)

func TestNewSet(t *testing.T) {
	t.Run("Creates an empty set", func(t *testing.T) {
		s := NewSet[int]()
		if s == nil {
			t.Fatalf("Set is nil")
		}

		if s.Size() != 0 {
			t.Fatalf("Set size is %d, expected %d", s.Size(), 0)
		}
	})

	t.Run("Creates a set with initial elements", func(t *testing.T) {
		s := NewSet[string]("1", "2", "3")
		if s == nil {
			t.Fatalf("Set is nil")
		}

		if s.Size() != 3 {
			t.Fatalf("Set size is %d, expected %d", s.Size(), 3)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Adds an element to an empty set", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(1)
		if s.Size() != 1 {
			t.Fatalf("Set size is %d, expected %d", s.Size(), 1)
		}

		if !s.Contains(1) {
			t.Fatalf("Set does not contain 1")
		}
	})

	t.Run("Adds an element to a non-empty set", func(t *testing.T) {
		s := NewSet[string]("one", "two", "three")
		s.Add("four")
		if s.Size() != 4 {
			t.Fatalf("Set size is %d, expected %d", s.Size(), 4)
		}

		if !s.Contains("four") {
			t.Fatalf(`Set does not contain "four"`)
		}
	})
}

func TestRemove(t *testing.T) {
	t.Run("Removes an element from a set", func(t *testing.T) {
		s := NewSet[int](1, 2, 3)
		s.Remove(2)
		if s.Size() != 2 {
			t.Fatalf("Set size is %d, expected %d", s.Size(), 2)
		}

		if s.Contains(2) {
			t.Fatalf("Set still contains 2")
		}
	})
}

func TestContains(t *testing.T) {
	t.Run("Returns true if an element is in the set", func(t *testing.T) {
		s := NewSet[int](1, 2, 3)
		if !s.Contains(2) {
			t.Fatalf("Set does not contain 2")
		}
	})

	t.Run("Returns false if an element is not in the set", func(t *testing.T) {
		s := NewSet[string]("one", "two", "three")
		if s.Contains("four") {
			t.Fatalf(`Set contains "four"`)
		}
	})
}

func TestSize(t *testing.T) {
	t.Run("Returns 0 for an empty set", func(t *testing.T) {
		s := NewSet[string]()
		if s.Size() != 0 {
			t.Fatalf("Set size is %d, expected %d", s.Size(), 0)
		}
	})

	t.Run("Returns the size of a non-empty set", func(t *testing.T) {
		s := NewSet[int](1, 2, 3)
		if s.Size() != 3 {
			t.Fatalf("Set size is %d, expected %d", s.Size(), 3)
		}
	})
}

func TestMembers(t *testing.T) {
	t.Run("Returns an empty slice for an empty set", func(t *testing.T) {
		s := NewSet[int]()
		members := s.Members()
		if len(members) != 0 {
			t.Fatalf("Set members length is %d, expected %d", len(members), 0)
		}
	})

	t.Run("Returns the members of a non-empty set", func(t *testing.T) {
		s := NewSet[string]("one", "two", "three")
		members := s.Members()
		if len(members) != 3 {
			t.Fatalf("Set members length is %d, expected %d", len(members), 3)
		}

		expected := []string{"one", "three", "two"}
		sort.Strings(members)
		if !slices.Equal(members, expected) {
			t.Fatalf(`Set members are %v, expected %v`, members, expected)
		}
	})
}
