package modulexp

import "testing"

func TestHelloSpencer(t *testing.T) {
	want := "Hello, Spencer!"
	if got := Hello("Spencer"); got != want {
		t.Errorf("Hello(\"Spencer\")=%q; want=%q", got, want)
	}
}

func TestVersion(t *testing.T) {
	want := "v1.0.0"
	if got := Version(); got != want {
		t.Errorf("Version()=%q; want=%q", got, want)
	}
}
