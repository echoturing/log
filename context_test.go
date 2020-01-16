package log

import (
	"context"
	"testing"
)

func TestKeyConflict(t *testing.T) {
	otherStructKey := struct{}{}
	var otherStructKeyInterface interface{} = otherStructKey
	ctx := NewContext(context.Background(), RequestIDWithUser{
		RequestID: "test request id",
		UserID:    "user_id",
	})
	if _, ok := ctx.Value(otherStructKey).(RequestIDWithUser); ok {
		t.Error("test key conflict error")
		return
	}
	if _, ok := ctx.Value(otherStructKeyInterface).(RequestIDWithUser); ok {
		t.Error("test key interface conflict error")
		return
	}
	Debug("test key conflict done!")
}
