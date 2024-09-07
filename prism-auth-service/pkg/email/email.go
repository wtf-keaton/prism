package email

import "net/mail"

func Validate(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
