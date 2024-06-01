package usersrv

import (
	"fmt"
	"http-server/internal/core/domain"
	"http-server/internal/core/ports"
)

type service struct {
	userRepository ports.UserRepository
}

func New(userRepository ports.UserRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) Login(name string, password string) (domain.User, error) {
	fmt.Println("srv Login")
	getUser, err := s.userRepository.GetByName(name)
	if err != nil {
		return domain.User{}, err
	}
	return getUser, nil
}

func (s *service) Create(createUser domain.User) (domain.User, error) {
	fmt.Println("srv Create")
	newUser, err := s.userRepository.Create(createUser)
	if err != nil {
		return domain.User{}, err
	}
	return newUser, nil
}

func (s *service) List() ([]domain.User, error) {
	fmt.Println("srv GetList")
	users, err := s.userRepository.List()
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (s *service) Get(id string) (domain.User, error) {
	fmt.Println("srv Get")
	user, err := s.userRepository.Get(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *service) Update(updateUser domain.User) (domain.User, error) {
	user, err := s.userRepository.Update(updateUser)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *service) Delete(id string) error {
	fmt.Println("srv Delete")
	err := s.userRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
