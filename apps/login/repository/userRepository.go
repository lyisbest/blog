package repository

import (
	"blog/apps/login/model"
	"blog/configuration"
	"golang.org/x/net/context"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUser(ctx context.Context, userName string, password string) (*model.User, error) {
	var user model.User
	result := configuration.DB.Where("username = ? AND password = ?", userName, password).Find(&user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUserName(ctx context.Context, userName string) (int, error) {
	var user model.User
	result := configuration.DB.Where("username=?", userName).Find(&user)
	if err := result.Error; err != nil {
		return 0, err
	}
	return int(result.RowsAffected), nil
}
