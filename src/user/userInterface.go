package user

import userEntity "service-code/model/entity/user"

type UserUsecase interface {
	Login(email, password string) (string, error)
	UserList() ([]*userEntity.User, error)
	UserById(id string) (*userEntity.User, error)
	CreateUser(fullname, email, password string) error
	UpdateUser(id, fullname, password string) error
	DeleteUser(id, requestorID string) error
}

type UserRepository interface {
	GetUserByEmailPassword(email string) (*userEntity.User, error)
	GetUserByEmail(email string) (*userEntity.User, error)
	GetListUsers() ([]*userEntity.User, error)
	GetUserByID(id string) (*userEntity.User, error)
	InsertUser(fullname, email, password string) error
	UpdateUser(id, fullname, password string) error
	DeleteUser(id string) error
}
