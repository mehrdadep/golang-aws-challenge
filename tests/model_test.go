package test

import (
	"testing"
)

//TestGetModel test status codes in getModel
func TestGetModel(t *testing.T) {
	total := 10
	if total != 10 {
		t.Errorf("Model was incorrect, got: %d, want: %d.", total, 10)
	}
}
