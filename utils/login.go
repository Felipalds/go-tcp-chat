package utils

import "go-tcp-chat/models"

func IsLoggedIn(currentUser *models.User) bool {
	if currentUser == nil || currentUser.Name == "" {
		return false
	}
	return true
}
