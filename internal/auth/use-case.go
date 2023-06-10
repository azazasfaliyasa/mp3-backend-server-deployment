package auth

import (
	"errors"
	"mini-project-3-azaz/internal/user"
	"mini-project-3-azaz/pkg/password"
)

type UseCase struct {
	R RepositoryInterface
}
type UseCaseInterface interface {
	Login(payload *Payload) (*ProfileItemWithToken, error)
	ShowProfile(tokenSigned string) (*user.User, error)
}

func NewUseCase(r RepositoryInterface) UseCaseInterface {
	return &UseCase{
		R: r,
	}
}

func (u *UseCase) Login(payload *Payload) (*ProfileItemWithToken, error) {
	user, _ := u.R.FindByUsername(payload.Username)
	if user.Username == "" {
		return nil, errors.New("user not found")
	}

	match := password.CheckPasswordHash(payload.Password, user.Password)
	if !match {
		return nil, errors.New("password is incorrect")
	}

	return u.R.Login(user)
}

func (u *UseCase) ShowProfile(tokenSigned string) (*user.User, error) {
	return u.R.ShowProfile(tokenSigned)
}
