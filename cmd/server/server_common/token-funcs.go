package server_common

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateNewToken() (string, time.Time) {
	now := time.Now()
	sha256Binary := sha256.Sum256([]byte(now.String()))
	token := hex.EncodeToString(sha256Binary[:])
	return token, now
}

func GetUserIdFromContext(ctx context.Context) (string, error) {
	v := ctx.Value(USER_ID_KEY)
	userId, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("user ID not found in the context")
	}

	return userId, nil
}
