package hash

import (
	"github.com/chenyusolar/framework/contracts/hash"
	"github.com/chenyusolar/framework/facades"
)

const (
	DriverBcrypt string = "bcrypt"
)

type Application struct {
}

func NewApplication() hash.Hash {
	driver := facades.Config.GetString("hashing.driver", "argon2id")

	if driver == DriverBcrypt {
		return NewBcrypt()
	}

	// Default set to argon2id
	return NewArgon2id()
}
