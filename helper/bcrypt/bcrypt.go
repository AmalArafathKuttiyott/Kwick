package bcrypt

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) string {
	hp, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		return ""
	}
	return string(hp)
}

func ComparePassword(up string, dp string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(up), []byte(dp))
	return err == nil
}
