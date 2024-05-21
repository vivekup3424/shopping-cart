package user

import (
	"database/sql"

	"github.com/vivekup3424/shopping-cart/types"
)

// following the repository pattern
type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}
func (s *Store) CreateUser(user types.User) (string, error) {
	var id string
	query := `
		INSERT INTO users(
		first_name,
		last_name, 
		email, 
		password)
		VALUES(
			$1, $2, $3, $4
		)
		RETURNING 
		id
		`
	err := s.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password).Scan(&id)
	if err != nil {
		return "-1", nil
	}
	return id, nil
}
func (s *Store) GetUserByID(id string) (*types.User, error) {
	var user types.User
	query := `
	SELECT id,
		first_name,
		last_name,
		email,
		password
	FROM users
	WHERE id = $1
	`
	err := s.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Store) GetUserByEmail(email string) (
	*types.User, error) {
	var user types.User
	query := `
		SELECT id,
			first_name,
			last_name,
			email,
			password
		FROM users
		WHERE email = $1
		`
	err := s.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (s *Store) UpdateUser(user types.User) error {
	query := `
	UPDATE users
	SET
	first_name = $1,
	last_name = $2,
	email = $3,
	password = $4
	WHERE id = $5
	`
	_, err := s.db.Exec(query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.ID)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) DeleteUser(id string) error {
	query := `
	DELETE FROM users
	WHERE id = $1
	`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
