package flexible_reversi_store

import (
	"fmt"
	"log"
	"slash_mochi/cmd/server/server_common"
	"time"
)

type FlexibleReversiStoreStructures struct {
	globalChatHistory []*ChatHistoryItem
	userInfos         []*UserInfo
}

type ChatHistoryItem struct {
	message string
	userId  string
}

func NewChatHistoryItem(message string, userId string) *ChatHistoryItem {
	return &ChatHistoryItem{
		message: message,
		userId:  userId,
	}
}

type ChatItem struct {
	Message  string
	UserName string
}

func NewChatItem(message string, userName string) *ChatItem {
	return &ChatItem{
		Message:  message,
		UserName: userName,
	}
}

type UserInfo struct {
	userId          string
	name            string
	expirationTimer time.Timer
}

func NewUserInfo(name string) *UserInfo {
	// generate new user ID
	userId, _ := server_common.GenerateNewToken()

	// start a timer for expiration
	expirationTimer := time.NewTimer(server_common.USER_EXPIRATION_DURATION)

	return &UserInfo{
		userId:          userId,
		name:            name,
		expirationTimer: *expirationTimer,
	}
}

func NewFlexibleReversiStructures() *FlexibleReversiStoreStructures {
	return &FlexibleReversiStoreStructures{
		globalChatHistory: make([]*ChatHistoryItem, 0),
		userInfos:         make([]*UserInfo, 0),
	}
}

// Methods

func (s *FlexibleReversiStoreStructures) AppendGlobalChatHistory(chat ChatHistoryItem) []*ChatHistoryItem {
	s.globalChatHistory = append(s.globalChatHistory, &chat)
	historyLength := len(s.globalChatHistory)

	// if chat history size is over the limit
	if historyLength > server_common.GLOBAL_CHAT_HISTORY_SIZE {
		// delete some oldest histories
		s.globalChatHistory = s.globalChatHistory[historyLength-server_common.GLOBAL_CHAT_HISTORY_SIZE:]
	}

	return s.globalChatHistory
}

// Returns user info according to a user ID string.
// If there is no userId in the user info slice, then returns error "invalid user ID"
func (s *FlexibleReversiStoreStructures) GetUserInfo(userId string) (*UserInfo, error) {
	for _, userInfo := range s.userInfos {
		if userId == userInfo.userId {
			return userInfo, nil
		}
	}

	log.Println("invalid user ID")
	return nil, fmt.Errorf("invalid user ID")
}

// Returns true if a userId is defined in the store, otherwise false.
func (s *FlexibleReversiStoreStructures) ValidateUserId(userId string) bool {
	for _, userInfo := range s.userInfos {
		if userId == userInfo.userId {
			return true
		}
	}

	log.Println("user ID not found")
	return false
}
