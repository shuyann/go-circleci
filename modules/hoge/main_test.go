package main

import "testing"

func TestEcho(t *testing.T) {
	if echo() != "hoge" {
		t.Errorf("echo() = %s but want %s", echo(), "hoge")
	}
}
