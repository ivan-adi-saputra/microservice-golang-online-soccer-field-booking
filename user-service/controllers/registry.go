package controllers

import (
	controllers "user-service/controllers/user"
	"user-service/services"
)

type registry struct {
	service services.ServiceRegistry
}

type ControllerRegistry interface {
	GetUserController() controllers.UserController
}

func NewControllerRegistry(service services.ServiceRegistry) ControllerRegistry {
	return &registry{service: service}
}

func (r *registry) GetUserController() controllers.UserController {
	return controllers.NewUserController(r.service)
}
