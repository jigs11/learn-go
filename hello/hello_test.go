package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jigs", "")
		want := "Hello, Jigs!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Rosie", "Spanish")
		want := "Hola, Rosie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello("Reggie", "French")
		want := "Bonjour, Reggie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in Sanskrit", func(t *testing.T) {
		got := Hello("Marcus", "Sanskrit")
		want := "Namaste, Marcus!"
		assertCorrectMessage(t, got, want)
	})

}
