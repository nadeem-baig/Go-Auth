package user

import (
	"database/sql"

	"github.com/nadeem-baig/go-auth/types"
	"github.com/nadeem-baig/go-auth/utils/logger"
)

type UserStore interface {
	GetUserByEmail(email string) (*types.User, error)
	CreateUser(user types.User) error
}

type Store struct {
	db *sql.DB
}

// Ensure Store implements UserStore
var _ UserStore = (*Store)(nil)

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail (email string) (*types.User, error)  {
	rows,err := s.db.Query("SELECT * FROM users WHERE email = ?",email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	u := new(types.User)
	for rows.Next() {
		u,err = scanRowIntoUser(rows)
		if err!= nil {
            return nil, err
        }
	}
	if u.ID == 0 {
		return nil, logger.Errorf("user not found")
	}
	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err!= nil {
        return nil, err
    }
	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil,nil
}

func (s *Store) CreateUser(user types.User) error {
	return nil
}