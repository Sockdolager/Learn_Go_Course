package mascot_test

import (
	"testing"

	"github.com/Sockdolager/Learn_Go_Course/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMoscot() != "Go Gopher" {
		t.Fatal("Wrong mascot :(")
	}
}
