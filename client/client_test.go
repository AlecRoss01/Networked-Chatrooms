package main

import (
	"testing"
)

func TestThirdPerson(t *testing.T) {
	want := "Alec likes dogs"
	if got := thirdPerson("likes dogs", "Alec"); got != want {
		t.Errorf("thirdPerson() = %q, want %q", got, want)
	}
}

func TestCoinFlip(t *testing.T) {
	want := true
	flip := coinFlip()
	var got bool
	if flip == "tails" || flip == "heads" {
		got = true
	} else {
		got = false
	}
	if got != want {
		t.Errorf("coinFlip() = %q, want heads or tails", flip)
	}
}

func TestDiceRoll(t *testing.T) {
	want := true
	roll := diceRoll()
	var got bool
	if roll >= 1 && roll <= 6 {
		got = true
	} else {
		got = false
	}
	if got != want {
		t.Errorf("diceRoll() = %q, want 1-6", roll)
	}
}

func TestEchoCommand(t *testing.T) {
	// The command the client would send
	msg := "!echo Echo this message!"

	// The expected server response
	want := "Echo this message!\n"

	got := commandHandler(msg)
	if got != want {
		t.Errorf("got = %q, want %q", got, want)
	}
}
