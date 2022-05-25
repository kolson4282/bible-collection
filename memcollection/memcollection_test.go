package memcollection_test

import (
	"testing"
)

func checkForError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func checkForNoError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("No Error Returned")
	}
}
