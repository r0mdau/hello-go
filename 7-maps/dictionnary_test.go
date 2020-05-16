package __maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("Dictionary v1", func(t *testing.T) {
		dict := map[string]string{"test": "definition of test word"}

		got := Search(dict, "test")
		want := "definition of test word"

		assertString(t, got, want)
	})

	Dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {

		got, _ := Dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := Dictionary.Search("coucou")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		Dictionary := Dictionary{}
		word := "flower"
		definition := "herbal plant"

		err := Dictionary.Add(word, definition)

		assertNoError(t, err)
		assertDefinition(t, Dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		Dictionary := Dictionary{word: definition}
		err := Dictionary.Add(word, "coucou")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, Dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	newDefinition := "new definition"

	t.Run("existing word", func(t *testing.T) {
		Dictionary := Dictionary{word: definition}
		err := Dictionary.Update(word, newDefinition)

		assertNoError(t, err)
		assertDefinition(t, Dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		Dictionary := Dictionary{}
		err := Dictionary.Update(word, newDefinition)

		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error")
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertDefinition(t *testing.T, Dictionary Dictionary, word string, definition string) {
	got, err := Dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word", err)
	}

	if got != definition {
		t.Errorf("got %q, want %q", got, definition)
	}
}
