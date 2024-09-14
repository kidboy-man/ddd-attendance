package helper

import (
	"fmt"
	"net/http"
	"strings"

	emailverifier "github.com/AfterShip/email-verifier"
	constant "github.com/kidboy-man/ddd-attendance/internal/domain/account/constants"
	generic "github.com/kidboy-man/ddd-attendance/internal/generics"
	"golang.org/x/crypto/bcrypt"
)

var (
	verifier = emailverifier.NewVerifier()
)

/**
 * CleanEmail cleans the provided email by removing any dots from the username and converting it to lowercase.
 * It also checks if the email format is valid. If the format is invalid, it returns an error.
 *
 * @param email string The email address to be cleaned.
 * @return cleanedEmail string The cleaned email address.
 * @return err error An error containing the HTTP status code, error code, and message if the email format is invalid.
 */
func CleanEmail(email string) (cleanedEmail string, err error) {
	res, err := verifier.Verify(email)
	if err != nil || !res.Syntax.Valid || strings.Contains(email, "+") {
		err = &generic.CustomError{
			Code:       constant.InvalidEmailFormatErrCode,
			HTTPStatus: http.StatusBadRequest,
			Message:    "Invalid email format",
		}
		return

	}

	tmp := strings.ToLower(strings.TrimSpace(email))
	slices := strings.Split(tmp, "@")
	username := strings.ReplaceAll(slices[0], ".", "")
	cleanedEmail = fmt.Sprintf("%s@%s", username, slices[1])
	return
}

/**
 * HashPassword hashes the provided password using bcrypt with a cost factor of 32.
 * It returns the hashed password and an error if there's an issue generating the hash.
 *
 * @param password string The password to be hashed.
 * @return hashed string The hashed password.
 * @return err error An error containing the bcrypt error if there's an issue generating the hash.
 */
func HashPassword(password string) (hashed string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return
	}

	hashed = string(bytes)
	return
}

/**
 * CheckPasswordHash checks if the provided password matches the hashed password.
 * It uses bcrypt's CompareHashAndPassword function to compare the provided password with the hashed password.
 * If the passwords match, it returns true; otherwise, it returns false.
 *
 * @param password string The password to be compared with the hashed password.
 * @param hash string The hashed password to be compared with the provided password.
 * @return isMatched bool A boolean value indicating whether the provided password matches the hashed password.
 */
func CheckPasswordHash(password, hash string) (isMatched bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	isMatched = err == nil
	return
}
