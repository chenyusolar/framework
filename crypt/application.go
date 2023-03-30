package crypt

import (
	"github.com/chenyusolar/framework/contracts/crypt"
)

type Application struct {
}

func NewApplication() crypt.Crypt {
	return NewAES()
}
