package iteration

import "testing"

// func TestRepeat(t *testing.T) {
// 	repeated := Repeat("a")
// 	expected := "aaaaa"

// 	if repeated != expected {
// 		t.Errorf("repeated is %q , expected is %q", repeated,expected)
// 	}
// }

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("ava")
	}
}