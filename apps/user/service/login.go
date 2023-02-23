package service

import (
	"blog/apps/user/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

type LoginService struct {
	userRepository repository.UserRepository
}

func NewLoginService(userRepository repository.UserRepository) *LoginService {
	return &LoginService{userRepository: userRepository}
}

func (s *LoginService) Login(ctx *gin.Context, userName string, password string) error {
	user, err := s.userRepository.GetUser(ctx, userName, password)
	if err != nil {
		log.Printf("GetUser failed, username: %v, password: %v, error: %v", userName, password, err)
		return err
	}
	if user == nil {
		err := errors.New("用户不存在")
		log.Printf("GetUser failed, error: %v", err)
		return err
	}
	ctx.SetCookie("user_cookie", user.Username, 1000, "/", "localhost", false, true)
	return nil
}
