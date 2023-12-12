package service

import (
	"bytes"
	"context"
	"devgym-http/src/errors"
	"devgym-http/src/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type Service struct {
	client  *http.Client
	urlBase string
}

func NewService(client *http.Client, url string) Service {
	return Service{
		client:  client,
		urlBase: url,
	}
}

func (s *Service) Create(ctx context.Context, deploy model.Deployment) (*model.Deployment, error) {
	j, err := json.Marshal(deploy)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/deployments", s.urlBase)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(j))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusBadRequest {
		return nil, errors.FromBadRequest(resp)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.FromHTTPResponse(resp)
	}

	createdDeploy := model.Deployment{}
	if err := json.NewDecoder(resp.Body).Decode(&createdDeploy); err != nil {
		return nil, err
	}

	return &createdDeploy, nil
}
