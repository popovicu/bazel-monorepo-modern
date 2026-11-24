package greeting

import "testing"

func TestFor(t *testing.T) {
	if got, want := For("Bazel"), "Hello, Bazel!"; got != want {
		t.Errorf("For(%q) = %q, want %q", "Bazel", got, want)
	}
}
