package scimsdk

import (
	"net/http"

	"github.com/strongdm/scimsdk/internal/api"
	"github.com/strongdm/scimsdk/internal/service"
)

type MockClient struct {
	apiInternalExecuteHTTPRequest func(*http.Request) (*http.Response, error)
	token                         string
}

func NewMockClient(apiInternalExecuteHTTPRequest func(*http.Request) (*http.Response, error), token string) Client {
	return &MockClient{apiInternalExecuteHTTPRequest, token}
}

func (mock *MockClient) Users() UserModule {
	mockApi := api.NewMockAPI(mock.apiInternalExecuteHTTPRequest)
	service := service.NewUserService(mockApi, mock.token)
	return &userModuleImpl{mock, service}
}

func (mock *MockClient) Groups() GroupModule {
	mockApi := api.NewMockAPI(mock.apiInternalExecuteHTTPRequest)
	service := service.NewGroupService(mockApi, mock.token)
	return &groupModuleImpl{mock, service}
}

func (mock *MockClient) GetProvidedURL() string {
	return ""
}
