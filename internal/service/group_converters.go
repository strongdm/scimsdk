package service

import (
	"encoding/json"
	"io"
)

func unmarshalGroupPageResponse(body io.ReadCloser) (*GroupPageResponse, error) {
	unmarshedResponse := &GroupPageResponse{}
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

func unmarshalGroupResponse(body io.ReadCloser) (*GroupResponse, error) {
	unmarshedResponse := &GroupResponse{}
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
