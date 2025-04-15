package internal

import "context"

type contextKey string

const UserIDKey contextKey = "userID"

func GetUserID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(UserIDKey).(string)
	return id, ok
}
