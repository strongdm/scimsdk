package module

import (
	"net/http"

	"github.com/strongdm/scimsdk/internal/api"
	"github.com/strongdm/scimsdk/internal/service"
)

func getMockedAPI(mockedFn func(*http.Request) (*http.Response, error)) api.API {
	return api.NewMockAPI(mockedFn)
}

func NewMockGroupModule(service service.GroupService) *groupModuleImpl {
	return &groupModuleImpl{service, ""}
}

func NewMockUserModule(svc service.UserService) *userModuleImpl {
	return &userModuleImpl{svc, ""}
}
