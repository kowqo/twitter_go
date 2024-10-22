package graphqlserver

import "context"

type AuthService interface {
	Register(ctx context.Context, input RegisterInput)
}

type RegisterInput struct {
	Email           string
	Username        string
	Password        string
	ConfirmPassword string
}

func (in *RegisterInput) Sanitize() {
	in.Email
}
