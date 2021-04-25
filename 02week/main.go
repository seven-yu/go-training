package main

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

// User user
type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

// GetUserByID get user by user_id
// If ErrNoRows appears, GetUserByID will return nil, nil
func GetUserByID(ctx context.Context, db *sql.DB, id uint) (*User, error) {
	user := &User{}
	err := db.QueryRowContext(ctx, "SELECT id,name, age FROM users WHERE id=?", id).Scan(&user.ID, &user.Name, &user.Age)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, errors.Wrap(err, "query users err")
	default:
		return user, nil
	}
}
