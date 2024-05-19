package userRepository

import (
	"database/sql"
	"errors"
	userEntity "service-code/model/entity/user"
	"service-code/src/user"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetListUsers() ([]*userEntity.User, error) {
	query := "SELECT id, fullname, email, password, created_at, updated_at FROM users;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*userEntity.User
	for rows.Next() {
		var user userEntity.User
		if err := rows.Scan(&user.ID, &user.Fullname, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetUserByID(id string) (*userEntity.User, error) {
	var user userEntity.User
	query := "SELECT id, fullname, email, password, created_at, updated_at FROM users WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Fullname, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmailPassword(email string) (*userEntity.User, error) {
	var user userEntity.User
	err := r.db.QueryRow("SELECT email, password FROM users WHERE email = $1", email).Scan(&user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*userEntity.User, error) {
	var user userEntity.User
	err := r.db.QueryRow("SELECT email FROM users WHERE email = $1", email).Scan(&user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) InsertUser(fullname, email, password string) error {
	_, err := r.db.Exec("INSERT INTO users (fullname, email, password) VALUES ($1, $2, $3)", fullname, email, password)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateUser(id, fullname, password string) error {
	query := "UPDATE users SET fullname = $2, password = $3 WHERE id = $1"
	_, err := r.db.Exec(query, id, fullname, password)
	if err == nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteUser(id string) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err == nil {
		return err
	}
	return nil
}
