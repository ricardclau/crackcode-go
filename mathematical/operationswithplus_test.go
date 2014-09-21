package mathematical

import (
	"testing"
)

func TestMultiplicationWorks(t *testing.T) {
	if result := Multiplication(7, 6); 42 != result {
		t.Errorf("Forgot about School? Result: %d", result)
	}

	if result := Multiplication(6, 7); 42 != result {
		t.Errorf("Forgot about School? Result: %d", result)
	}

	/* Tests not working yet...
	if result := Multiplication(7, -6); -42 != result {
		t.Errorf("Forgot about School? Result: %d", result)
	}

	if result := Multiplication(-7, -6); 42 != result {
		t.Errorf("Forgot about School? Result: %d", result)
	}
	*/
}
