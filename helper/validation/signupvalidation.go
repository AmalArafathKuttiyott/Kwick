package validation

import (
	models "kwick/model/request"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrorString string
var Valid bool

func ValidateSignupDetails(u models.RequestBody) (map[string]string, bool) {
	// Validating First name
	ErrorString, Valid = ValidateFirstName(u.UserFirstName)
	if !Valid {
		return map[string]string{"Error": ErrorString}, Valid
	}
	// Validating Phone number
	ErrorString, Valid = ValidatePhone(u.UserPhone)
	if !Valid {
		return map[string]string{"Error": ErrorString}, Valid
	}
	// Validating Email
	ErrorString, Valid = ValidateEmail(u.UserEmail)
	if !Valid {
		return map[string]string{"Error": ErrorString}, Valid
	}
	// Validating Password
	Valid = ValidatePassword(u.UserPassword)
	if !Valid {
		return map[string]string{"Error": "Enter a strong password including  uppercase, lowercase, numbers, special character"}, Valid
	}
	// Comparing password and confirm password
	match := ComparePasswords(u.UserPassword, u.UserConfirmPassword)
	if !match {
		return map[string]string{"Error": "Password does not match"}, match
	}
	return map[string]string{"Status": "Validation Successful"}, true
}
func ValidateFirstName(f string) (string, bool) {
	if f == "" {
		return "FirstName is missing", false
	}
	if len(f) < 3 {
		return "Invalid name. Please enter a valid name.", false
	}
	validNameRegex := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !validNameRegex.MatchString(f) {
		return "Invalid name. The first name should only contain alphabetical characters.", false
	}
	return "", true
}
func ValidatePhone(p string) (string, bool) {
	if p == "" {
		return "Phone number is missing", false
	}
	if len(p) != 10 {
		return "Invalid phone number. The phone number should be exactly 10 digits long", false
	}
	validPhoneRegex := regexp.MustCompile(`^[0-9]+$`)
	if !validPhoneRegex.MatchString(p) {
		return "Invalid phone number. The phone number should only contain numeric digits.", false
	}
	return "", true
}
func ValidateEmail(e string) (string, bool) {
	if e == "" {
		return "Email, which is required, is missing", false
	}
	if strings.TrimSpace(e) != e {
		return "Email should not contain leading or trailing whitespace", false
	}
	parts := strings.Split(e, "@")
	domain := parts[len(parts)-1]
	if strings.Contains(domain, "..") {
		return "Email contains consecutive dots in the domain", false
	}
	for _, r := range e {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '.' && r != '_' && r != '-' && r != '@' {
			return "Email contains invalid characters", false
		}
	}
	if strings.Count(e, "@") > 1 {
		return "Email contains multiple @ symbols", false
	}
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	val := regexp.MustCompile(regex).MatchString(e)
	if !val {
		return "Email is not valid. Please use a correct email format", false
	}
	if utf8.RuneCountInString(e) != len(e) {
		return "Email should not contain international characters", false
	}
	return "", true
}
func ValidatePassword(p string) bool {
	if len(p) < 8 {
		return false
	}
	if !containsUppercase(p) {
		return false
	}
	if !containsLowercase(p) {
		return false
	}
	if !containsDigit(p) {
		return false
	}
	if !containsSpecialCharacter(p) {
		return false
	}
	return true
}
func containsUppercase(s string) bool {
	regExp := regexp.MustCompile(`[A-Z]`)
	return regExp.MatchString(s)
}
func containsLowercase(s string) bool {
	regExp := regexp.MustCompile(`[a-z]`)
	return regExp.MatchString(s)
}
func containsDigit(s string) bool {
	regExp := regexp.MustCompile(`\d`)
	return regExp.MatchString(s)
}
func containsSpecialCharacter(s string) bool {
	regExp := regexp.MustCompile(`[^a-zA-Z0-9]`)
	return regExp.MatchString(s)
}
func ComparePasswords(p string, cp string) bool {
	return p == cp
}
