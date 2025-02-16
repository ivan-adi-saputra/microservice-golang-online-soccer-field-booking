package repositories

import (
	repositories "user-service/repositories/user"

	"gorm.io/gorm"
)

type registry struct {
	DB *gorm.DB
}

type RepositoryRegistry interface {
	GetUser() repositories.UserRepository
}

func NewRepositoryRegistry(db *gorm.DB) *registry {
	return &registry{DB: db}
}

func (r *registry) GetUser() repositories.UserRepository {
	return repositories.NewUserRepository(r.DB)
}
