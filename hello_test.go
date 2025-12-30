package main

import (
	"testing"
)

func TestHello(t *testing.T){
	// got :=hello("Chris")
	// want := "hello, Chris"

	// if got != want {
	// 	t.Errorf("got %q , want %q",got,want)
	// }

	t.Run("say hello to person",func(t *testing.T){
			got :=hello("Chris")
	want := "hello, Chris"

	// assertCorrectMessage(t,got,want)
	assertCorrectMessage(t, got, want)
	})

	t.Run("saying just 'hello, World' when empty string is provided",func(t *testing.T){
		got := hello("")
		want := "hello, Worlds"
	assertCorrectMessage(t, got, want)	
	})
}
func assertCorrectMessage (t testing.TB,got, want string ){
	t.Helper()
		if got!= want {
			t.Errorf("got %q, want %q"  , got  , want)
		}

}
