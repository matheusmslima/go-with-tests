package strututil

import (
	"fmt"
	"testing"
)

func TestShout(t *testing.T) {
	got := Shout("Hello")
	want := "HELLO"
	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}

func ExampleShout() {
	shout := Shout("hello")
	fmt.Println(shout)
	//Output: HELLO
}
