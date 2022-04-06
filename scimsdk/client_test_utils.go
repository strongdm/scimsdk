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

func NewMockClient(apiInternalExecuteHTTPRequest func(*http.Request) (*http.Response, error), token string) IClient {
	return &MockClient{apiInternalExecuteHTTPRequest, token}
}

func (mock *MockClient) Users() IUserModule {
	mockApi := api.NewAPI()
	mockApi.(*api.API).SetInternalExecuteHTTPRequest(mock.apiInternalExecuteHTTPRequest)
	service := service.NewUserService(mockApi, mock.token)
	return &UserModule{mock, service}
}

func (mock *MockClient) Groups() IGroupModule {
	mockApi := api.NewAPI()
	mockApi.(*api.API).SetInternalExecuteHTTPRequest(mock.apiInternalExecuteHTTPRequest)
	service := service.NewGroupService(mockApi, mock.token)
	return &GroupModule{mock, service}
}

func (mock *MockClient) GetProvidedURL() string {
	return ""
}
