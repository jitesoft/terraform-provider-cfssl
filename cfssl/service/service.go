package service

import (
	"github.com/jitesoft/terraform-provider-cfssl/cfssl/client"
)

type Service struct {
	client *client.Client
}

func New(client *client.Client) *Service {
	service := Service{}
	service.client = client

	return &service
}
