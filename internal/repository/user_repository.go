package repository

import (
	"database/sql"
	"user-management/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return []models.User{}, nil
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, name, email FROM users WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	return r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
}

func (r *UserRepository) DeleteUser(id int) error {
	return r.db.QueryRow("DELETE FROM users WHERE id = $1", id).Err()
}
