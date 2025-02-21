package user

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/huuloc2026/restfulapi-gorm.git/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	if db == nil {
		panic("NewStore: received nil database connection")
	}
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {

	rows, err := s.db.Query("SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE email = $1", email)

	if err != nil {
		fmt.Println("Query error:", err) // ✅ Log lỗi truy vấn
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		user, err := scanRowIntoUser(rows)
		if err != nil {
			fmt.Println("Error scanning row:", err) // ✅ Log lỗi khi scan dữ liệu
			return nil, err
		}
		// fmt.Printf("User found: %+v\n", user) // ✅ Log user tìm thấy
		return user, nil
	}

	//fmt.Println("User not found") // ✅ Log nếu không tìm thấy user
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
	rows, err := s.db.Query("SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE id = ?", id)
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
func (s *Store) CreateUser(user types.User) error {

	_, err := s.db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return fmt.Errorf("user with email %s already exists", user.Email)
		}
		return err
	}
	return nil
}
