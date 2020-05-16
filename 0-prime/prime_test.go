package __prime

import (
	"strconv"
	"testing"
)

func Testprime(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("1 is prime", func(t *testing.T) {
		got := strconv.FormatBool(isPrime(1))
		want := "true"

		assertCorrectMessage(t, got, want)
	})

	t.Run("2 is not prime", func(t *testing.T) {
		got := strconv.FormatBool(isPrime(2))
		want := "false"

		assertCorrectMessage(t, got, want)
	})

	t.Run("97 is prime", func(t *testing.T) {
		got := strconv.FormatBool(isPrime(2))
		want := "true"

		assertCorrectMessage(t, got, want)
	})

	t.Run("2351 is not prime", func(t *testing.T) {
		got := strconv.FormatBool(isPrime(2))
		want := "false"

		assertCorrectMessage(t, got, want)
	})

	t.Run("resultToPrint return primeResultPrint const if prime", func(t *testing.T) {
		got := resultToPrint(true)
		want := primeResultPrint

		assertCorrectMessage(t, got, want)
	})

	t.Run("resultToPrint return notPrimeResultPrint const if not prime", func(t *testing.T) {
		got := resultToPrint(false)
		want := notPrimeResultPrint

		assertCorrectMessage(t, got, want)
	})
}
