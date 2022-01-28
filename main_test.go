package main

import "testing"

func TestGoWithTests(t *testing.T) {
	if GoWithTests() != "Go with Tests" {
		t.Errorf("Not again!!!")
	}
}
