package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestUpdateMessage(t *testing.T) {
	expected := "Hello, Go"

	updateMessage(expected)

	if msg != expected {
		t.Errorf("expected %s, but got %s", expected, msg)
	}
}

func TestPrintMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	updateMessage("Hello, universe")
	printMessage()

	w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe") {
		t.Errorf("Expected 'Hello, universe', but got it is no there")
	}
}

func TestMain(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	expected := []string{"universe", "cosmos", "world"}
	for _, word := range expected {
		if !strings.Contains(output, word) {
			t.Errorf("Expected universe, cosmos or world, but got it is no there")
		}
	}
}
