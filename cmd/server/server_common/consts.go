package server_common

import "time"

const (
	GLOBAL_CHAT_HISTORY_SIZE      int = 200
	STORE_INTERFACE_CHANNELS_SIZE int = 100

	// can-i-use-up-food
	MAX_PROJECT_PERIOD time.Duration = 30 * 24 * time.Hour // 30 days

	// flexible-reversi
	USER_EXPIRATION_DURATION time.Duration = 30 * time.Minute
	USER_ID_KEY              string        = "user_id"
)
