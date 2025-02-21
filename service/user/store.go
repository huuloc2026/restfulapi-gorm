package user

import (
	"database/sql"
	"fmt"
	"github.com/huuloc2026/restfulapi-gorm.git/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		user, err := scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, fmt.Errorf("user not found")
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(&user.ID, &user.Name, &user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (s *Store) CreateUser(user types.User) error {
	return nil
}
