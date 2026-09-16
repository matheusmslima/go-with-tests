package somemodule

import "testing"

func TestHello(t *testing.T) {
	got := BuySandals()
	want := "Havaianas"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
