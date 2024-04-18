package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"unicode"

	"github.com/Eiphoria/GoReversi/internal/repository"
)

var ErrInvalidData = errors.New("invalid data")

type Service struct {
	repo *repository.Repository
	salt string
}

func New(rep *repository.Repository, salt string) *Service {
	return &Service{
		repo: rep,
		salt: salt,
	}
}

func (s *Service) CreateUser(ctx context.Context, username, password string) error {
	if usernameLen := len(username); usernameLen < 3 || usernameLen > 20 {
		return ErrInvalidData
	}

	if passwordLen := len(password); passwordLen < 3 || passwordLen > 20 {
		return ErrInvalidData
	}

	var hasDigit, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasDigit && !hasSpecial {
		return ErrInvalidData
	}

	hash := getMD5Hash(password, s.salt)

	err := s.repo.CreateUser(ctx, username, hash)
	if err != nil {
		return fmt.Errorf("repo create user: %w", err)
	}

	return nil
}

func getMD5Hash(text, salt string) string {
	salted := text + salt
	hash := md5.Sum([]byte(salted))
	return hex.EncodeToString(hash[:])
}
