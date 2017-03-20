package helper

import "testing"

func TestIsValidPermision(t *testing.T) {
	if IsValidPermision("private") != true {
		t.Error("Se esperaba true y se obtuvo false")
	}
	if IsValidPermision("friends") != true {
		t.Error("Se esperaba true y se obtuvo false")
	}
	if IsValidPermision("public") != true {
		t.Error("Se esperaba true y se obtuvo false")
	}
}
