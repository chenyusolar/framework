package cache

import (
	"github.com/chenyusolar/framework/facades"
)

func prefix() string {
	return facades.Config.GetString("cache.prefix") + ":"
}
