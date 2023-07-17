package validation

import (
	models "kwick/model/request"
)

func SignupFormCheck(u models.RequestBody) (string, bool) {
	if u.UserFirstName == "" {
		return "First name required", false
	}
	if u.UserEmail == "" {
		return "Email required", false
	}
	if u.UserPhone == "" {
		return "Phone required", false
	}
	if u.UserPassword == "" {
		return "Password required", false
	}
	if u.UserConfirmPassword == "" {
		return "Confirm password required", false
	}
	return "Validation Successfull", true
}
