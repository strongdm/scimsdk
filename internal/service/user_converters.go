package service

import (
	"encoding/json"
	"io"
)

func unmarshalUserPageResponse(body io.ReadCloser) (*UserPageResponse, error) {
	unmarshedResponse := &UserPageResponse{}
	buff, err := convertResponseBodyToBytes(body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buff, &unmarshedResponse)
	if err != nil {
		return nil, err
	}
	return unmarshedResponse, nil
}

func unmarshalUserResponse(body io.ReadCloser) (*UserResponse, error) {
	unmarshedResponse := &UserResponse{}
	buff, err := convertResponseBodyToBytes(body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buff, &unmarshedResponse)
	if err != nil {
		return nil, err
	}
	return unmarshedResponse, nil
}
