package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world, from Victor"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
