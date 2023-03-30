package filesystem

import (
	"github.com/chenyusolar/framework/facades"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.Storage = NewStorage()
}

func (database *ServiceProvider) Boot() {

}
