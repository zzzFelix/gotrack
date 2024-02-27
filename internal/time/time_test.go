package time

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	t.Run("should parse time when given hour and minute in 24h format", func(t *testing.T) {
		// given
		input := "16:45"
		expected := time.Date(0, 1, 1, 16, 45, 0, 0, time.UTC)

		// when
		actual := ParseTime(input)

		// then
		if expected != actual {
			t.Fatalf("unexpected result. expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("should parse time when only given hour in 24h format", func(t *testing.T) {
		// given
		input := "16"
		expected := time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC)

		// when
		actual := ParseTime(input)

		// then
		if expected != actual {
			t.Fatalf("unexpected result. expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("should return time.Now() when input cannot be parsed", func(t *testing.T) {
		// given
		input := "not a time"
		expected := time.Now()

		// when
		actual := ParseTime(input)

		// then
		difference := actual.Sub(expected)
		if difference > 1*time.Second {
			t.Fatalf("unexpected result. expected: %v, actual: %v", expected, actual)
		}
	})
}
