package starter_test

import (
	"testing"

	starter "github.com/anjali29singh/testing_go"
)

// unit testing
func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("anjali")
	expected := "Welcom anjali singh!"
	if expected != greeting {
		t.Errorf("failed test")
	}
}
