package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	a := 8
	b := 4

	got := Add(a, b)
	want := 12

	if got != want {
		t.Errorf("Got: %q, want: %q", got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
	
}
