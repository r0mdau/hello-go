package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("greetingPrefix returns in english if not recognized", func(t *testing.T) {
		got := greetingPrefix("Greek")
		want := "Hello, "

		assertCorrectMessage(t, got, want)
	})

	t.Run("greetingPrefix returns in spanish", func(t *testing.T) {
		got := greetingPrefix("Spanish")
		want := "Hola, "

		assertCorrectMessage(t, got, want)
	})

	t.Run("greetingPrefix returns in french", func(t *testing.T) {
		got := greetingPrefix("French")
		want := "Bonjour, "

		assertCorrectMessage(t, got, want)
	})
}
