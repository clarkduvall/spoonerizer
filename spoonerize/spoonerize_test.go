package spoonerize

import (
	"testing"
)

func assertEqual(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("expected \"%v\" got \"%v\"", expected, actual)
	}
}

func assertSpoonerized(t *testing.T, s1 string, s2 string) {
	assertEqual(t, string(Spoonerize([]byte(s1))), s2)
}

func TestSpoonerize(t *testing.T) {
	assertSpoonerized(t, "", "")
	assertSpoonerized(t, " ", " ")
	assertSpoonerized(t, "hello world", "wello horld")
	assertSpoonerized(t, "Hello World", "Wello Horld")
	assertSpoonerized(t, "Hello world", "Wello horld")
	assertSpoonerized(t, "hello World", "wello Horld")
	assertSpoonerized(t, "shoot food", "foot shood")
	assertSpoonerized(t, "captain crunch", "craptain cunch")
	assertSpoonerized(t, "scrunchy flapper", "flunchy scrapper")
	assertSpoonerized(t, "foo", "foo")
	assertSpoonerized(t, "f b", "b f")
	assertSpoonerized(t, "hacker news", "nacker hews")
	assertSpoonerized(t, "abc def ghi jkl", "abc ghef di jkl")
	assertSpoonerized(t, "xoop bip quick jim", "boop xip jick quim")
	assertSpoonerized(t, "hello the world", "wello the horld")
	assertSpoonerized(t, "bob neither jim", "job neither bim")
	assertSpoonerized(t, "hello \"world\"", "wello \"horld\"")
	assertSpoonerized(t, "bad \u2013 character", "chad \u2013 baracter")
	assertSpoonerized(t, "bad \u2013character", "bad \u2013character")
	assertSpoonerized(t, "bad c\u2013haracter", "cad b\u2013haracter")
}
