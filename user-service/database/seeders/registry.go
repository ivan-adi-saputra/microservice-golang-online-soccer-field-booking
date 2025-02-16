package seeders

import "gorm.io/gorm"

type Registry struct {
	DB *gorm.DB
}

type SeederRegistry interface {
	Run()
}

func NewSeederRegistry(db *gorm.DB) *Registry {
	return &Registry{DB: db}
}

func (s *Registry) Run() {
	RunRoleSeeder(s.DB)
	RunUserSeeder(s.DB)
}
