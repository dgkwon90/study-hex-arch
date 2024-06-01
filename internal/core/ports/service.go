package ports

import "http-server/internal/core/domain"

type UserService interface {
	Login(name string, password string) (domain.User, error)
	Create(domain.User) (domain.User, error)
	Get(id string) (domain.User, error)
	List() ([]domain.User, error)
	Update(domain.User) (domain.User, error)
	Delete(id string) error
}
