package validation

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/application/dto"
	"gopkg.in/validator.v2"
)

func ValidateUser(userInput *dto.UserInput) error {
	if err := validator.Validate(userInput); err != nil {
		return err
	}

	return nil
}
