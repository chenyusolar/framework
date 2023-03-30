package validation

import (
	consolecontract "github.com/chenyusolar/framework/contracts/console"
	"github.com/chenyusolar/framework/facades"
	"github.com/chenyusolar/framework/validation/console"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.Validation = NewValidation()
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]consolecontract.Command{
		&console.RuleMakeCommand{},
	})
}
