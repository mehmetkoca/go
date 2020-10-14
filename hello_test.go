package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mehmet")
	want := "Hello, Mehmet"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
