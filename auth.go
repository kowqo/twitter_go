package graphqlserver

import (
	"context"
	"fmt"
	"regexp"
	"strings"
)

const (
	UsernameMinLength = 2
	PasswordMinLength = 6
)

var emailRegex = regexp.MustCompile("^[\\w-.]+@([\\w-]+\\.)+[\\w-]{2,7}$")

type AuthService interface {
	Register(ctx context.Context, input RegisterInput) (AuthResponse, error)
}

type AuthResponse struct {
	AccessToken string
	User        User
}

type RegisterInput struct {
	Email           string
	Username        string
	Password        string
	ConfirmPassword string
}

func (in *RegisterInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)

	in.Username = strings.TrimSpace(in.Username)
}

func (in *RegisterInput) Validate() error {
	if len(in.Username) < UsernameMinLength {
		return fmt.Errorf("%w: username too short, (%d) character min length", ErrValidation, UsernameMinLength)
	}

	if len(in.Password) < PasswordMinLength {
		return fmt.Errorf("%w: password too short", ErrValidation)
	}

	if !emailRegex.MatchString(in.Email) {
		return fmt.Errorf("%w: Invalid email", ErrValidation)
	}

	if in.Password != in.ConfirmPassword {
		return fmt.Errorf("%w: Invalid password", ErrValidation)
	}

	return nil
}
