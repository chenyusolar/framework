package mock

import (
	mocksaccess "github.com/chenyusolar/framework/contracts/auth/access/mocks"
	mocksauth "github.com/chenyusolar/framework/contracts/auth/mocks"
	mockscache "github.com/chenyusolar/framework/contracts/cache/mocks"
	mocksconfig "github.com/chenyusolar/framework/contracts/config/mocks"
	mocksconsole "github.com/chenyusolar/framework/contracts/console/mocks"
	mocksorm "github.com/chenyusolar/framework/contracts/database/orm/mocks"
	mocksevent "github.com/chenyusolar/framework/contracts/event/mocks"
	mocksfilesystem "github.com/chenyusolar/framework/contracts/filesystem/mocks"
	mocksgrpc "github.com/chenyusolar/framework/contracts/grpc/mocks"
	mocksmail "github.com/chenyusolar/framework/contracts/mail/mocks"
	mocksqueue "github.com/chenyusolar/framework/contracts/queue/mocks"
	mocksvalidate "github.com/chenyusolar/framework/contracts/validation/mocks"
	"github.com/chenyusolar/framework/facades"
	"github.com/chenyusolar/framework/log"
)

func Cache() *mockscache.Store {
	mockCache := &mockscache.Store{}
	facades.Cache = mockCache

	return mockCache
}

func Config() *mocksconfig.Config {
	mockConfig := &mocksconfig.Config{}
	facades.Config = mockConfig

	return mockConfig
}

func Artisan() *mocksconsole.Artisan {
	mockArtisan := &mocksconsole.Artisan{}
	facades.Artisan = mockArtisan

	return mockArtisan
}

func Orm() (*mocksorm.Orm, *mocksorm.Query, *mocksorm.Transaction, *mocksorm.Association) {
	mockOrm := &mocksorm.Orm{}
	facades.Orm = mockOrm

	return mockOrm, &mocksorm.Query{}, &mocksorm.Transaction{}, &mocksorm.Association{}
}

func Event() (*mocksevent.Instance, *mocksevent.Task) {
	mockEvent := &mocksevent.Instance{}
	facades.Event = mockEvent

	return mockEvent, &mocksevent.Task{}
}

func Log() {
	facades.Log = log.NewApplication(log.NewTestWriter())
}

func Mail() *mocksmail.Mail {
	mockMail := &mocksmail.Mail{}
	facades.Mail = mockMail

	return mockMail
}

func Queue() (*mocksqueue.Queue, *mocksqueue.Task) {
	mockQueue := &mocksqueue.Queue{}
	facades.Queue = mockQueue

	return mockQueue, &mocksqueue.Task{}
}

func Storage() (*mocksfilesystem.Storage, *mocksfilesystem.Driver, *mocksfilesystem.File) {
	mockStorage := &mocksfilesystem.Storage{}
	mockDriver := &mocksfilesystem.Driver{}
	mockFile := &mocksfilesystem.File{}
	facades.Storage = mockStorage

	return mockStorage, mockDriver, mockFile
}

func Validation() (*mocksvalidate.Validation, *mocksvalidate.Validator, *mocksvalidate.Errors) {
	mockValidation := &mocksvalidate.Validation{}
	mockValidator := &mocksvalidate.Validator{}
	mockErrors := &mocksvalidate.Errors{}
	facades.Validation = mockValidation

	return mockValidation, mockValidator, mockErrors
}

func Auth() *mocksauth.Auth {
	mockAuth := &mocksauth.Auth{}
	facades.Auth = mockAuth

	return mockAuth
}

func Gate() *mocksaccess.Gate {
	mockGate := &mocksaccess.Gate{}
	facades.Gate = mockGate

	return mockGate
}

func Grpc() *mocksgrpc.Grpc {
	mockGrpc := &mocksgrpc.Grpc{}
	facades.Grpc = mockGrpc

	return mockGrpc
}
