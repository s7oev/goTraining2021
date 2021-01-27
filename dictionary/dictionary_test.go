package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := dictionary {
		"something": "this is something",
	}

	result, err := dictionary.Search("something")
	expect := "this is something"

	if err != nil {
		t.Errorf("Did not expect an error but got %s", err)
	}

	if result != expect {
		t.Errorf("Got %s but expected %s", result, expect)
	}
}

func TestSearchNotFound(t *testing.T) {
	dictionary := dictionary {
		"something": "this is something",
	}

	_, err := dictionary.Search("notsomething")

	if err != ErrNotFound {
		t.Error("Expected error word not found")
	}
}

func TestAddNonExistingWord(t *testing.T) {
	dictionary := dictionary{}
	key := "somethingNonExisting"
	value := "definition of somethingNonExisting"
	dictionary.Add(key, value)

	result, err := dictionary.Search(key)
	wanted := value

	if err != nil {
		t.Errorf("Did not expect an error but got %s", err)
	}

	if result != wanted {
		t.Errorf("Expected %s but got %s", wanted, result)
	}
}

func TestAddExistingWord(t *testing.T) {
	key := "something"
	value := "this is something"

	dictionary := dictionary {
		key: value,
	}

	gotErr := dictionary.Add(key, "this is new something")

	if gotErr != ErrWordExists {
		t.Error("Expected an error that word exists")
	}

	gotDefinition, _ := dictionary.Search(key)

	if gotDefinition != value {
		t.Errorf("Expected definition %s but got definition %s", value, gotDefinition)
	}
}

func TestDeleteExistingWord(t *testing.T) {
	key := "something"
	value := "this is something"

	dictionary := dictionary {
		key: value,
	}

	got := dictionary.Delete(key)

	if got != nil {
		t.Errorf("Did not expect an error but got the error %s", got)
	}

	_, got = dictionary.Search(key)

	if got == nil {
		t.Error("Expected an error for non-existing word but did not get one")
	}
}
