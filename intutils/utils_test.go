package intutils

import "testing"

func TestAbs(t *testing.T) {
	t.Run("Returns the absolute value of a positive number", func(t *testing.T) {
		if Abs(7) != 7 {
			t.Fatalf("Abs(7) is %d, expected 7", Abs(7))
		}
	})

	t.Run("Returns the absolute value of a negative number", func(t *testing.T) {
		if Abs(-7) != 7 {
			t.Fatalf("Abs(-7) is %d, expected 7", Abs(-7))
		}
	})
}

func TestAtoi(t *testing.T) {
	t.Run("Converts a string to a positive integer", func(t *testing.T) {
		if Atoi("123") != 123 {
			t.Fatalf(`Atoi("123") is %d, expected 123`, Atoi("123"))
		}
	})

	t.Run("Converts a string to a negative integer", func(t *testing.T) {
		if Atoi("-456") != -456 {
			t.Fatalf(`Atoi("-456") is %d, expected -456`, Atoi("-456"))
		}
	})

	t.Run("Panics if the string cannot be converted to an integer", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf(`Atoi("one hundred twenty three") did not panic`)
			}
		}()
		Atoi("one hundred twenty three")
	})
}

func TestPow(t *testing.T) {
	t.Run("Returns the power of a positive integer", func(t *testing.T) {
		if Pow(2, 3) != 8 {
			t.Fatalf("Pow(2, 3) is %d, expected 8", Pow(2, 3))
		}
	})

	t.Run("Returns the power of a negative integer", func(t *testing.T) {
		if Pow(-2, 3) != -8 {
			t.Fatalf("Pow(-2, 3) is %d, expected -8", Pow(-2, 3))
		}
	})

	t.Run("Returns the power of an integer raised to zero", func(t *testing.T) {
		if Pow(2, 0) != 1 {
			t.Fatalf("Pow(2, 0) is %d, expected 1", Pow(2, 0))
		}
	})
}
