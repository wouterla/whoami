package main

import (
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {
	expected := "8080"
	out := GetPort()
	if out != expected {
		t.Errorf("%s expected; got %s", expected, out)
	}

	/* Set env variable PORT to 1234 and handle error if any */
	err := os.Setenv("PORT", "1234")
	if err != nil {
		t.Fatal(err)
	}

	expected = "1234"
	out = GetPort()
	if out != expected {
		t.Errorf("%s expected; got %s", expected, out)
	}
}
