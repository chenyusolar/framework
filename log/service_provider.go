package log

import "github.com/chenyusolar/framework/facades"

type ServiceProvider struct {
}

func (log *ServiceProvider) Register() {
	facades.Log = NewLogrusApplication()
}

func (log *ServiceProvider) Boot() {

}
