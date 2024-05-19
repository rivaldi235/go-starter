package userUsecase

import (
	"errors"
	"fmt"
	userEntity "service-code/model/entity/user"
	"service-code/pkg/middleware"
	"service-code/pkg/validation"
	"service-code/src/user"

	"golang.org/x/crypto/bcrypt"
)

type userUC struct {
	userRepo user.UserRepository
}

func NewUserUsecase(userRepo user.UserRepository) user.UserUsecase {
	return &userUC{userRepo}
}

func (uc *userUC) Login(email, password string) (string, error) {
	user, err := uc.userRepo.GetUserByEmailPassword(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateTokenJwt(user.Email, 3)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *userUC) UserList() ([]*userEntity.User, error) {
	users, err := uc.userRepo.GetListUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *userUC) UserById(id string) (*userEntity.User, error) {
	user, err := uc.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUC) CreateUser(fullname, email, password string) error {
	user, _ := uc.userRepo.GetUserByEmail(email)

	if user != nil {
		return errors.New("email already exists")
	}

	if !validation.ValidatePasswordFormat(password) {
		return errors.New("the password must contain at least 1 uppercase letter, 1 lowercase letter, 1 digit, 1 special character, and must be at least 8 characters long")
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	fmt.Println(fullname, email, hashedPassword)

	if err := uc.userRepo.InsertUser(fullname, email, hashedPassword); err != nil {
		return err
	}

	return nil
}

func (uc *userUC) UpdateUser(id, fullname, password string) error {
	_, err := uc.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}

	if !validation.ValidatePasswordFormat(password) {
		return errors.New("the password must contain at least 1 uppercase letter, 1 lowercase letter, 1 digit, 1 special character, and must be at least 8 characters long")
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	if err := uc.userRepo.UpdateUser(id, fullname, hashedPassword); err != nil {
		return err
	}

	return nil
}

func (uc *userUC) DeleteUser(id, requestorID string) error {
	_, err := uc.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}

	if id == requestorID {
		return errors.New("cannot delete yourself")
	}

	if err := uc.userRepo.DeleteUser(id); err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
