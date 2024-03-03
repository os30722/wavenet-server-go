package utils

import (
	"context"
	"errors"
)

type userContextKey string

const UidKey = "UID"
const ClaimsKey = "claims"

func GetUserContextKey(str string) userContextKey {
	return userContextKey(str)
}

func GetUid(ctx context.Context) (int, error) {
	claims, ok := ctx.Value(GetUserContextKey(ClaimsKey)).(map[string]any)
	if !ok {
		return -1, errors.New("No uid found")
	}

	uid := int(claims[UidKey].(int))
	return uid, nil
}
