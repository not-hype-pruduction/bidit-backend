package grpc

import (
	"google.golang.org/grpc"
)

// ServiceRegistrar умеет регистрировать себя на gRPC сервере
type ServiceRegistrar interface {
	Register(server *grpc.Server)
}

// Registry хранит все gRPC хендлеры
type Registry struct {
	services []ServiceRegistrar
}

// NewRegistry создает новый реестр хендлеров
func NewRegistry(services ...ServiceRegistrar) *Registry {
	return &Registry{
		services: services,
	}
}

// RegisterAll регистрирует все сервисы на gRPC сервере
func (r *Registry) RegisterAll(server *grpc.Server) {
	for _, service := range r.services {
		service.Register(server)
	}
}
