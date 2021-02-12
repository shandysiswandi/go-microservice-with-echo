package util_test

import (
	"go-rest-echo/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword_Check_Five_Character_From_Front_No_Error(t *testing.T) {
	is := assert.New(t)

	actual, err := util.HashPassword("password")
	expected := "$2a$10"

	is.Nil(err)
	is.Equal(expected, actual[0:6])
}

func TestCheckPasswordHash(t *testing.T) {
	is := assert.New(t)

	actual := util.CheckPasswordHash("password", "$2a$10$mjqqoczR7odoHg/npdnwcuJCk4GHUDYTrkX48vuy/tNq7P/V/wAGi")
	expected := true

	is.Equal(expected, actual)
}
