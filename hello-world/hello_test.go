package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("", "")

		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodia", "Spanish")
		want := "Hola, Elodia"

		assertCorrectMessage(t, got, want)

	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodia", "French")
		want := "Bonjur, Elodia"

		assertCorrectMessage(t, got, want)

	})
}
