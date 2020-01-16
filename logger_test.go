package log

import "testing"

func TestDebugWithContext(t *testing.T) {
	ctx := NewDefaultContext()
	DebugWithContext(ctx, "test debug", "hello", "world")
}
