package usecase

import (
	"github.com/b4ckslash/rental-app/services/user/entity"
	"github.com/b4ckslash/rental-app/services/user/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(email, password, role string) error
	Login(email, password string) (*entity.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (uc *userUsecase) Register(email, password, role string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &entity.User{
		Email:        email,
		PasswordHash: string(hash),
		Role:         role,
	}
	return uc.repo.Create(user)
}

func (uc *userUsecase) Login(email, password string) (*entity.User, error) {
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}
