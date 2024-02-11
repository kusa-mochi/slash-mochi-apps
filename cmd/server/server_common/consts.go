package server_common

import "time"

const (
	GLOBAL_CHAT_HISTORY_SIZE      int           = 200
	STORE_INTERFACE_CHANNELS_SIZE int           = 100
	USER_EXPIRATION_DURATION      time.Duration = 30 * time.Minute
	USER_ID_KEY                   string        = "user_id"
)
