package services

import (
	"user-service/repositories"
	services "user-service/services/user"
)

type registry struct {
	repository repositories.RepositoryRegistry
}

type ServiceRegistry interface {
	GetUser() services.UserService
}

func NewServiceRegistry(repository repositories.RepositoryRegistry) ServiceRegistry {
	return &registry{repository: repository}
}

func (r *registry) GetUser() services.UserService {
	return services.NewUserService(r.repository)
}
