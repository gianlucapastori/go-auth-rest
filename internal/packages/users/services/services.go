package services

import (
	"errors"
	"fmt"
	"net/smtp"

	"github.com/gianlucapastori/nausicaa/cfg"
	"github.com/gianlucapastori/nausicaa/internal/entities"
	"github.com/gianlucapastori/nausicaa/internal/packages/users"
	"github.com/gianlucapastori/nausicaa/pkg/jwt"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type userService struct {
	repo  users.Repo
	cfg   *cfg.Config
	sugar *zap.SugaredLogger
}

func New(repo users.Repo, cfg *cfg.Config, sugar *zap.SugaredLogger) users.Services {
	return &userService{repo: repo, cfg: cfg, sugar: sugar}
}

func (uS *userService) Register(user *entities.User) (*entities.User, error) {
	if err := uS.isEmailAvailable(user.Email); err != nil {
		return nil, err
	}

	if err := uS.isUsernameAvailable(user.Username); err != nil {
		return nil, err
	}

	hash, err := user.HashPassword()
	if err != nil {
		return nil, err
	}

	user.Password = hash

	u, err := uS.repo.InsertUser(user)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uS *userService) Login(user *entities.User, pwd string) error {
	if err := user.ComparePassword(pwd); err != nil {
		return err
	}

	return nil
}

func (uS *userService) SendPwdResetEmail(user *entities.User) error {
	// use hashed pwd as payload, you can request as much as you  want, but just one will do
	token, err := jwt.RequestPwdToken(user.Password, uS.cfg)
	if err != nil {
		return err
	}

	email := []string{user.Email}

	auth := smtp.PlainAuth("", uS.cfg.SERVER.EMAIL, uS.cfg.SERVER.EMAIL_PWD, "smtp.gmail.com")

	subject := "Subject: new password request\r\n"
	body := fmt.Sprintf("http://127.0.0.1:8080/api/v1/users/change-password?token=%s\r\nBye\r\n", token)

	msg := subject + "\r\n" + body

	err = smtp.SendMail("smtp.gmail.com:587", auth, uS.cfg.SERVER.EMAIL, email, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func (uS *userService) ChangePasswordByEmail(email string, pwd string) error {
	u, err := uS.FetchByEmail(email)
	if err != nil {
		return err
	}

	if u == nil {
		return errors.New("users does not exists? wtf")
	}

	u.Password = pwd

	hash, err := u.HashPassword()
	if err != nil {
		return err
	}

	return uS.repo.ChangeUserPasswordByEmail(email, hash)
}

func (uS *userService) FetchByEmail(email string) (*entities.User, error) {
	return uS.repo.FetchUserByEmail(email)
}

func (uS *userService) FetchByUsername(username string) (*entities.User, error) {
	return uS.repo.FetchUserByUsername(username)
}

func (uS *userService) isEmailAvailable(email string) error {
	u, err := uS.FetchByEmail(email)
	if err != nil {
		uS.sugar.Errorf(err.Error())
		return err
	}
	if u.Id != uuid.Nil {
		return errors.New("email already in use")
	}
	return nil
}

func (uS *userService) isUsernameAvailable(username string) error {
	u, err := uS.FetchByUsername(username)
	if err != nil {
		uS.sugar.Errorf(err.Error())
		return err
	}
	if u.Id != uuid.Nil {
		return errors.New("username already in use")
	}
	return nil
}
