package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	expect := 4
	if got != expect {
		t.Errorf("Expected %d but got %d", expect, got)
	}
}

func ExampleAdd() {
	result := Add(2, 2)
	fmt.Println(result)
	// Output: 5
}
