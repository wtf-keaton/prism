package methods

import (
	"context"
	"errors"
	"fmt"
	"prism-auth-service/internal/postgres"
	"prism-auth-service/pkg/email"
	"prism-auth-service/pkg/passwordHash"
	"time"
)

const ( // role
	Default = iota
	HR
	Moderator
)

func SignUp(login, password string) error {

	if !email.Validate(login) {
		return errors.New("login is not email")
	}

	tx, err := postgres.Get().Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	var exists bool
	err = tx.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", login).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check user existence: %v", err)
	}

	if exists {
		return postgres.ErrUserExists
	}

	passwordHashed, err := passwordHash.New(password)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(),
		"INSERT INTO users(email, password_hash, role, created_at) VALUES($1, $2, $3, $4)",
		login, passwordHashed, Default, time.Now())

	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
