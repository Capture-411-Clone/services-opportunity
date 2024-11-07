package jwd

import (
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	jwt.StandardClaims
	Roles []int `json:"roles"`
}

func (c *CustomClaims) GetID() int {
	id, err := strconv.Atoi(c.Id)
	if err != nil {
		panic("cannot convert jwt claim id to int")
	}
	return id
}
