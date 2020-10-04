package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("Add new key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"

		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("Add existing key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"

		dictionary = Dictionary{key: value}

		err := dictionary.Add(key, "new value")

		assertError(t, err, ErrKeyExist)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		newValue := "new update"
		dictionary := Dictionary{key: value}

		err := dictionary.Update(key, newValue)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newValue)
	})

	t.Run("new key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(key, value)

		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	value := "this is just a test"
	dictionary := Dictionary{key: value}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)

	assertError(t, err, ErrNotFound)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatalf("key %s should exist. err: %v", "test", err)
	}

	if got != value {
		t.Errorf("got %q want %q", got, value)
	}
}
