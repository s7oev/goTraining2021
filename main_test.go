package main

import "testing"

func assert(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %s but wanted %s", got, want)
	}
}

func assertTranslatedGreeting (t *testing.T, name, language, expect string) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()
		got := Greeting(name, language)
		want := expect
		assert(t, got, want)
	}
}

func TestGreeting(t *testing.T) {
	given := "Lars"
	got := Greeting(given, "")
	want := "Hello Lars"
	assert(t, got, want)
}

func TestGreetingEmpty(t *testing.T) {
	given := ""
	got := Greeting(given, "")
	want := "Hello World"
	assert(t, got, want)
}

func TestGreetingWithLanguages(t *testing.T) {
	t.Run("bulgarian",
		assertTranslatedGreeting(t, "Lars", Bulgarian, "здравей Lars"))

	t.Run("romanian",
		assertTranslatedGreeting(t, "Lars", Romanian, "Salut Lars"))
}