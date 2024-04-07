package service

import (
	"context"
	"errors"
	"fmt"
	"unicode"

	"github.com/Eiphoria/GoReversi/internal/repository"
)

var ErrInvalidData = errors.New("invalid data")

type Service struct {
	repo *repository.Repository
}

func New(rep *repository.Repository) *Service {
	return &Service{
		repo: rep,
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

	err := s.repo.CreateUser(ctx, username, password)
	if err != nil {
		return fmt.Errorf("repo create user: %w", err)
	}

	return nil
}
