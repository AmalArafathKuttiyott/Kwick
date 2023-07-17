package validation

import (
	request "kwick/model/request"
)

func SigninFormCheck(u request.RequestBody) (string, bool) {
	if u.UserSigninData == "" {
		return "Email or Phone required", false
	}
	if u.UserPassword == "" {
		return "Password required", false
	}
	return "", true
}
func ValidateSigninDetails(u request.RequestBody) (string, bool) {
	if len(u.UserSigninData) > 10 {
		ErrorString, Valid = ValidateEmail(u.UserSigninData)
		if !Valid {
			return ErrorString, Valid
		}
	} else {
		ErrorString, Valid = ValidatePhone(u.UserSigninData)
		if !Valid {
			return ErrorString, Valid
		}
	}
	return "Validation Successful", true
}
