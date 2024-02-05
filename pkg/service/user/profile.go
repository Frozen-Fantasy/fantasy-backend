package user

import (
	"github.com/Frozen-Fantasy/fantasy-backend.git/pkg/models/user"
	"github.com/google/uuid"
)

func (s *Service) GetUserInfo(userID uuid.UUID) (user.UserInfoModel, error) {
	userInfo, err := s.storage.GetUserInfo(userID)
	if err != nil {
		return userInfo, err
	}
	return userInfo, nil
}
