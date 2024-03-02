package utils

import (
	"context"
)

type userContextKey string

const UidKey = "UID"
const ClaimsKey = "claims"

func GetUserContextKey(str string) userContextKey {
	return userContextKey(str)
}

func GetUid(ctx context.Context) (int, bool) {
	claims, ok := ctx.Value(GetUserContextKey(ClaimsKey)).(map[string]any)
	if !ok {
		return -1, ok
	}

	uid := int(claims[UidKey].(int))
	return uid, true
}
