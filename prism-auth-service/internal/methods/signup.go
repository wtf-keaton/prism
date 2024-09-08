package methods

import (
	"context"
	"errors"
	"fmt"
	"prism-auth-service/internal/database"
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

	tx, err := database.Get().Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	var exists bool
	err = tx.QueryRow(context.Background(), "select exists(select 1 from users where email=$1)", login).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check user existence: %v", err)
	}

	if exists {
		return errors.New("user exists")
	}

	passwordHashed, err := passwordHash.New(password)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(),
		"insert into users(email, password_hash, role, created_at) values($1, $2, $3, $4)",
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
