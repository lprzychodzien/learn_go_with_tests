package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("t")
	expected := "ttttt"

	if got != expected {
		t.Errorf("expected %q got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
