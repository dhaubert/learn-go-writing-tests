package main

import (
	"testing"
	"os"
)

func BenchMarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Douglas", "es_ES");
	}
}

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() //it helps to test suite alert about original line and not this one
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to a person", func (t *testing.T) {
		got := Hello("Douglas", "")
		want := "Hello, Douglas"
		assertCorrectMessage(t, got, want);
	})

	t.Run("saying hello to a person in Spanish", func (t *testing.T) {
		got := Hello("Douglas", "es_ES")
		want := "Hola, Douglas"
		assertCorrectMessage(t, got, want);
	})

	t.Run("saying hello to a person in French", func (t *testing.T) {
		got := Hello("Douglas", "fr_FR")
		want := "Bonjour, Douglas"
		assertCorrectMessage(t, got, want);
	})


	t.Run("saying hello to the world when an empty string is supplied", func (t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want);
	})
	
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}