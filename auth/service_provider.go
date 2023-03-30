package auth

import (
	"context"

	"github.com/chenyusolar/framework/auth/access"
	"github.com/chenyusolar/framework/auth/console"
	contractconsole "github.com/chenyusolar/framework/contracts/console"
	"github.com/chenyusolar/framework/facades"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.Auth = NewAuth(facades.Config.GetString("auth.defaults.guard"))
	facades.Gate = access.NewGate(context.Background())
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]contractconsole.Command{
		&console.JwtSecretCommand{},
		&console.PolicyMakeCommand{},
	})
}
