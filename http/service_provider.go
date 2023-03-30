package http

import (
	consolecontract "github.com/chenyusolar/framework/contracts/console"
	"github.com/chenyusolar/framework/facades"
	"github.com/chenyusolar/framework/http/console"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.RateLimiter = NewRateLimiter()
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]consolecontract.Command{
		&console.RequestMakeCommand{},
		&console.ControllerMakeCommand{},
		&console.MiddlewareMakeCommand{},
	})
}
