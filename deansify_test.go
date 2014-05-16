package deansify

import "testing"

func TestStripAnsi(t *testing.T) {
	before := `[1m[36m (0.1ms)[0m  [1mBEGIN[0m`
	expected := " (0.1ms)  BEGIN"
	actual := stripAnsi(before)

	if actual != expected {
		t.Errorf("Ansi strip failed: expected %s, got  %s", expected, actual)
	}
}
