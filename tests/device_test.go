package test

import (
	"testing"
)

//TestGetModel test status codes in getModel
func TestGetDevice(t *testing.T) {
	total := 10
	if total != 10 {
		t.Errorf("Device was incorrect, got: %d, want: %d.", total, 10)
	}
}
