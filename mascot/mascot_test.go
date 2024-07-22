package mascot_test

import (
	"testing"

	"sooz.com/go-demo-1/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong Mascot :(")
	}
}
