package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestCopyWord(t *testing.T) {
	got := CopyWordRepeat("Matheus", 1)
	want := "Matheus"
	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func BenchmarkCopyWord(b *testing.B) {
	for b.Loop() {
		CopyWordRepeat("Restaurant", 5)
	}
}

func ExampleRepeat() {
	sum := Repeat("a", 5)
	fmt.Println(sum)
	//Output: aaaaa
}

func ExampleCopyWordRepeat() {
	word := CopyWordRepeat("Restaurant", 5)
	fmt.Println(word)
	//Output: Restaurant
}
