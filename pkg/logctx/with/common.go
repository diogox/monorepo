package with

import "go.uber.org/zap"

const keyUserID = "user_id"

func UserID(userID string) Field {
	return func() zap.Field {
		return zap.String(keyUserID, userID)
	}
}
