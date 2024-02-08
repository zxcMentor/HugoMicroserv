package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"user/internal/models"
)

type UserRepository interface {
	CreateUser(email, hashepassword string) error
	ProfileUser(id int32) (*models.UserDTO, error)
	ListUsers() ([]*models.UserDTO, error)
}

type UserRepo struct {
	Postgres *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) CreateUser(email, hashepassword string) error {
	query := `INSERT INTO users (email, hashepassword) VALUES ($1, $2)`
	result, err := u.Postgres.Exec(query, email, hashepassword)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	// Получение количества затронутых строк
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %v", err)
	}

	return nil

}

func (u *UserRepo) ProfileUser(id int32) (*models.UserDTO, error) {
	var user *models.UserDTO
	query := `SELECT * FROM users WHERE id = $1` //посмотреть как вернуть юзера только с id & email
	err := u.Postgres.Get(&user, query, id)
	if err != nil {
		log.Println("err user not exist")
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) ListUsers() ([]*models.UserDTO, error) {
	var users []*models.UserDTO
	query := `SELECT * FROM users ` //тоже самое как и выше
	err := u.Postgres.Get(&users, query)
	if err != nil {
		log.Println("err dont get users")
		return nil, err
	}
	return users, nil
}
