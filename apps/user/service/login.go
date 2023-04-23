package service

import (
	"blog/apps/user/repository"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type LoginService struct {
	userRepository *repository.UserRepository
}

func NewLoginService(userRepository *repository.UserRepository) *LoginService {
	return &LoginService{userRepository: userRepository}
}

func (s *LoginService) Login(ctx *gin.Context, userName string, password string) error {
	_, err := s.userRepository.GetUser(ctx, userName, password)
	if err != nil {
		log.Printf("GetUser failed, username: %v, password: %v, error: %v", userName, password, err)
		return err
	}
	token, err := utils.GenerateToken(userName)
	if err != nil {
		log.Printf("token generate failed, error: %v", err)
		return err
	}
	ctx.Header("token", token)
	return nil
}
