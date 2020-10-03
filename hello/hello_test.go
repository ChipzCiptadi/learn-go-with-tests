package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	}

	t.Run("Hello with names argument", func(t *testing.T) {
		got := Hello("Chipz", "")
		want := "Hello, Chipz!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello with no argument", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello in Bahasa", func(t *testing.T) {
		got := Hello("Chipz", "Bahasa")
		want := "Halo, Chipz!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello in Batak", func(t *testing.T) {
		got := Hello("Chipz", "Batak")
		want := "Horas, Chipz!"

		assertCorrectMessage(t, got, want)
	})
}
