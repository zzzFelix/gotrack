package database

import "testing"

func TestReplaceHomePath(t *testing.T) {
	t.Run("should replace ~ with home path", func(t *testing.T) {
		// given
		path := "~/whatever"

		// when
		actual, err := replaceHomePath(path)

		// then
		if err != nil {
			t.Fatalf("an error occurred, %v", err)
		}
		if actual == path { // result is non-deterministic but path should have changed
			t.Fatalf("could not replace tilde")
		}
	})

	t.Run("should leave non-home path untouched", func(t *testing.T) {
		// given
		path := "not/a/home/directory"

		// when
		actual, err := replaceHomePath(path)

		// then
		if err != nil {
			t.Fatalf("an error occurred, %v", err)
		}
		if actual != path {
			t.Fatalf("path must not change, new path: %s", actual)
		}
	})
}
