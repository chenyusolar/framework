package crypt

import (
	"github.com/chenyusolar/framework/facades"
)

type ServiceProvider struct {
}

func (crypt *ServiceProvider) Register() {
	facades.Crypt = NewApplication()
}

func (crypt *ServiceProvider) Boot() {

}
