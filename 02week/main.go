package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

// User user
type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

var (
	ErrNotFound = errors.New("not found")
	ErrInternal = errors.New("internal err")
)

// GetUserByID get user by user_id
// If ErrNoRows appears, GetUserByID will return nil, nil
func GetUserByID(ctx context.Context, db *sql.DB, id uint) (*User, error) {
	user := &User{}
	err := db.QueryRowContext(ctx, "SELECT id,name, age FROM users WHERE id=?", id).Scan(&user.ID, &user.Name, &user.Age)

	switch {
	case err == sql.ErrNoRows:
		return nil, errors.Wrapf(ErrNotFound, "exec SELECT id,name, age FROM users WHERE id=%s err: %+v", id, err)
	case err != nil:
		return nil, errors.Wrapf(ErrInternal, "exec SELECT id,name, age FROM users WHERE id=%s err: %+v", id, err)
	default:
		return user, nil
	}
}

func BizCode() {
	user, err := GetUserByID(context.Background(), &sql.DB{}, 100)

	if errors.Is(err, ErrNotFound) {
		// do something
	}

	if errors.Is(err, ErrInternal) {
		// do something
	}

	// do something
	fmt.Sprintln(user)
}
