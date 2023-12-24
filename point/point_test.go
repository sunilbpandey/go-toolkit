package point

import "testing"

func TestPoint(t *testing.T) {
	t.Run("Returns the distance between two points", func(t *testing.T) {
		p := NewPoint(1, 2)
		q := NewPoint(4, 6)
		if p.Distance(q) != 7 {
			t.Fatalf("p.Distance(q) is %d, expected 7", p.Distance(q))
		}
	})

	t.Run("Returns the distance between two points with negative coordinates", func(t *testing.T) {
		p := NewPoint(-1, 2)
		q := NewPoint(-4, 6)
		if p.Distance(q) != 7 {
			t.Fatalf("p.Distance(q) is %d, expected 7", p.Distance(q))
		}
	})
}
