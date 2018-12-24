package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Miro")
	want := "Hello Miro"

	if got != want {
		t.Errorf("got '%s' 9want '%s'", got, want)
	}
}
