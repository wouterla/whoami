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

	/* Reset PORT */
	err = os.Setenv("PORT", "")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBindAddr(t *testing.T) {
	expected := ":8080"
	out := BindAddr()

	if out != expected {
		t.Errorf("%s expected; got %s", expected, out)
	}
}

func TestWhoamis(t *testing.T) {
	expected := []Whoami{
		{"FOO", "foo"},
		{"BAR", "bar"},
		{"LOL", "lol"},
	}

	/* Set environment variables */
	for _, whoami := range expected {
		key := "WHOAMI_" + whoami.Key
		err := os.Setenv(key, whoami.Value)
		if err != nil {
			t.Fatal(err)
		}
	}

	out := GetWhoamis()
	lout, lexpected := len(out), len(expected)
	if lout != lexpected {
		t.Errorf("%d elements expected; got %d", lexpected, lout)
	}

	for i := 0; i < lout; i++ {
		o, e := out[i], expected[i]
		if o.Key != e.Key || o.Value != e.Value {
			t.Errorf("%v expected; got %v", e, o)
		}
	}
}

func TestWhoamiFromEnvStr(t *testing.T) {
	expected := &Whoami{"FOO", "foo"}
	out := WhoamiFromEnvStr("WHOAMI_FOO=foo")
	if out.Key != expected.Key {
		t.Errorf("%v expected; got %v", expected.Key, out.Key)
	}

	if out.Value != expected.Value {
		t.Errorf("%v expected; got %v", expected.Value, out.Value)
	}
}
