package log

import (
	"context"
	"encoding/json"

	"github.com/segmentio/ksuid"
)

type Key struct {
}

var (
	// 用interface包了以后，即便是空的struct相比较也不一样了。。。
	key interface{} = Key{}
)

type RequestIDWithUser struct {
	RequestID string      `json:"requestId"`
	UserID    json.Number `json:"userId"`
}

func NewContext(parent context.Context, riu RequestIDWithUser) context.Context {
	return context.WithValue(parent, key, riu)
}

func NewDefaultContext() context.Context {
	return context.WithValue(context.Background(), key, RequestIDWithUser{
		RequestID: ksuid.New().String(),
	})
}

func FromContext(ctx context.Context) RequestIDWithUser {
	requestIDUser, ok := ctx.Value(key).(RequestIDWithUser)
	if !ok {
		requestIDUser.RequestID = ksuid.New().String()
	}
	return requestIDUser
}

func NewFromContext(parent context.Context) context.Context {
	ctx := context.Background()
	return NewContext(ctx, FromContext(parent))
}
