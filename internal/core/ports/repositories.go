package ports

import "http-server/internal/core/domain"

type UserRepository interface {
	Create(domain.User) (domain.User, error)
	GetByName(name string) (domain.User, error)
	Get(id string) (domain.User, error)
	List() ([]domain.User, error)
	Update(domain.User) (domain.User, error)
	Delete(id string) error
}
