package main

import "testing"

func TestEcho(t *testing.T) {
	if echo() != "foo" {
		t.Errorf("echo() = %s but want %s", echo(), "foo")
	}
}
