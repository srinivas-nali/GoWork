package gowork

import "testing"

func TestOne(t *testing.T) {
	expected := "Hello"
	actual := printValue("Hello")
	if expected != actual {
		t.Fail()
	}
}
