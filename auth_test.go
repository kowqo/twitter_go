package graphqlserver

import (
	"errors"
	"testing"
)

func TestRegisterInput_Sanitize(t *testing.T) {

	input := RegisterInput{
		Username:        " bob ",
		Email:           " bob@example.com ",
		Password:        "password",
		ConfirmPassword: "password",
	}

	want := RegisterInput{
		Username:        "bob",
		Email:           "bob@example.com",
		Password:        "password",
		ConfirmPassword: "password",
	}

	input.Sanitize()

	if input != want {
		t.Errorf("Sanitize() got: %v want: %v", input, want)
	}
}

func TestRegisterInput_Validate(t *testing.T) {
	testCases := []struct {
		name  string
		input RegisterInput
		err   error
	}{
		{
			name: "valid",
			input: RegisterInput{
				Username:        "bob",
				Email:           "bob@example.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: nil,
		},
		{
			name: "invalid name",
			input: RegisterInput{
				Username:        "2",
				Email:           "bob@example.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		},
		{
			name: "invalid email",
			input: RegisterInput{
				Username:        "2asd",
				Email:           "bobexample.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		},
		{
			name: "invalid conf",
			input: RegisterInput{
				Username:        "2asd",
				Email:           "bob@example.com",
				Password:        "pass2word",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()

			if tc.err != nil {
				//
				if !errors.Is(err, tc.err) {
					t.Errorf("it s error")
				}
			} else {
				//
			}
		})
	}
}
