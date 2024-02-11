package utils // Custom validator for phone number
import "github.com/go-playground/validator/v10"

// TODO: Add regex validation for indian number
func ValidatePhone(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()
	if len(phoneNumber) != 0 {
		// Check if the phone number starts with '+91' and has 10 digits
		if len(phoneNumber) == 13 && phoneNumber[:3] == "+91" {
			return true
		}
		return false
	}
	return true // Allow empty phone number
}
