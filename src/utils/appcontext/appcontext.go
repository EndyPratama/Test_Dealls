package appcontext

import (
	"context"
	"time"
)

type contextKey string

const (
	requestId        contextKey = "RequestId"
	userAgent        contextKey = "UserAgent"
	userAgentID      contextKey = "UserAgentID"
	requestStartTime contextKey = "RequestStartTime"
)

func SetRequestId(ctx context.Context, rid string) context.Context {
	return context.WithValue(ctx, requestId, rid)
}

func GetRequestId(ctx context.Context) string {
	rid, ok := ctx.Value(requestId).(string)
	if !ok {
		return ""
	}
	return rid
}

func SetUserAgent(ctx context.Context, ua string) context.Context {
	return context.WithValue(ctx, userAgent, ua)
}

func GetUserAgent(ctx context.Context) string {
	ua, ok := ctx.Value(userAgent).(string)
	if !ok {
		return ""
	}
	return ua
}

func SetRequestStartTime(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, requestStartTime, t)
}

func GetRequestStartTime(ctx context.Context) time.Time {
	t, _ := ctx.Value(requestStartTime).(time.Time)
	return t
}

func SetUserIDAgent(ctx context.Context, uia string) context.Context {
	return context.WithValue(ctx, userAgentID, uia)
}

func GetUserIDAgent(ctx context.Context) string {
	ua, ok := ctx.Value(userAgentID).(string)
	if !ok {
		return ""
	}
	return ua
}
