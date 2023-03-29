package main

import "testing"

func TestThirdPerson(t *testing.T) {
	want := "Alec likes dogs"
	if got := thirdPerson("likes dogs", "Alec"); got != want{
		t.Errorf("thirdPerson() = %q, want %q", got, want)
	}
}