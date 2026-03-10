package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello world"

	wg.Add(2)
	go updateMessage("x")
	go updateMessage("Goodbye world")
	wg.Wait()

	if msg != "Goodbye world" {
		t.Errorf("incorrect value in msg")
	}
}
