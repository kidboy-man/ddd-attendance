package helper

import (
	"net/http"
	"testing"

	constant "github.com/kidboy-man/ddd-attendance/internal/domain/account/constants"
	generic "github.com/kidboy-man/ddd-attendance/internal/generics"
	"github.com/stretchr/testify/assert"
)

func TestCleanEmail_InvalidEmailFormat(t *testing.T) {
	invalidEmail := "invalid.email"
	cleanedEmail, err := CleanEmail(invalidEmail)
	assert.Error(t, err)
	assert.Equal(t, constant.InvalidEmailFormatErrCode, err.(*generic.CustomError).Code)
	assert.Equal(t, http.StatusBadRequest, err.(*generic.CustomError).HTTPStatus)
	assert.Equal(t, "Invalid email format", err.(*generic.CustomError).Message)
	assert.Equal(t, "", cleanedEmail)
}

func TestCleanEmail_InvalidEmailFormatMultipleAtSymbols(t *testing.T) {
	invalidEmail := "invalid.email@email.com@invalid.com"
	cleanedEmail, err := CleanEmail(invalidEmail)
	assert.Error(t, err)
	assert.Equal(t, constant.InvalidEmailFormatErrCode, err.(*generic.CustomError).Code)
	assert.Equal(t, http.StatusBadRequest, err.(*generic.CustomError).HTTPStatus)
	assert.Equal(t, "Invalid email format", err.(*generic.CustomError).Message)
	assert.Equal(t, "", cleanedEmail)
}

func TestCleanEmail_ValidEmailFormat(t *testing.T) {
	validEmail := "valid.email@example.com"
	cleanedEmail, err := CleanEmail(validEmail)
	assert.NoError(t, err)
	assert.Equal(t, "validemail@example.com", cleanedEmail)
}

func TestHashPassword_SuccessHash(t *testing.T) {

	password := "1234zxcV!"
	hashed, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEqual(t, "", hashed)
}

func TestCheckPasswordHash_HashDoesNotMatchPassword(t *testing.T) {
	password := "1234zxcV=!"
	hash := "$2a$14$My9Uam0fyoij3LlsSWwj1uSVGxb4BTbYKwXirqGRsuHPeBIZd45ZK"
	isMatched := CheckPasswordHash(password, hash)
	assert.False(t, isMatched)
}

func TestCheckPasswordHash_PasswordMatched(t *testing.T) {
	password := "1234zxcV!"
	hash := "$2a$14$My9Uam0fyoij3LlsSWwj1uSVGxb4BTbYKwXirqGRsuHPeBIZd45ZK"
	isMatched := CheckPasswordHash(password, hash)
	assert.True(t, isMatched)
}
