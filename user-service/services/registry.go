package services

import (
	"user-service/repositories"
	services "user-service/services/user"
)

type registry struct {
	repository repositories.RepositoryRegistry
}

type Registry interface {
	GetUser() services.UserService
}

func NewServiceRegistry(repository repositories.RepositoryRegistry) Registry {
	return &registry{repository: repository}
}

func (r *registry) GetUser() services.UserService {
	return services.NewUserService(r.repository)
}
