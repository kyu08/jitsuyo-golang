package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrUnknown        = errors.New("unknown error")
)

func main() {
	userService := userService{}
	if err := userService.Authenticate(); err != nil {
		log.Fatal(err)
	}
}

type userService struct {
	userRepo userRepository
}

func (s *userService) Authenticate() error {
	if err := s.userRepo.FindUser("hoge"); err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			log.Print("ユーザーが存在しなかった場合のハンドリングをする")
		} else {
			return fmt.Errorf("failed to userRepo.FindUser: %w", err)
		}
	}
	return nil
}

type userRepository struct {
	db dbHandler
}

func (r *userRepository) FindUser(id string) error {
	if err := r.db.Query("SELECT * FROM users WHERE id = ?", id); err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}
	return nil
}

type dbHandler struct{}

func (h *dbHandler) Query(sql string, args ...interface{}) error {
	return ErrRecordNotFound
	// return ErrUnknown
}
