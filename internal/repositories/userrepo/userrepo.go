package userrepo

import (
	"encoding/json"
	"http-server/internal/core/domain"
	"http-server/pkg/apperrors"
	"http-server/pkg/uidgen"

	"github.com/matiasvarela/errors"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *memkvs {
	return &memkvs{kvs: map[string][]byte{}}
}

func (repo *memkvs) Create(createUser domain.User) (domain.User, error) {
	bytes, err := json.Marshal(createUser)
	if err != nil {
		return domain.User{}, errors.New(apperrors.InvalidInput, err, "user fails at marshal into json string", "")
	}

	id := uidgen.New().New()
	repo.kvs[id] = bytes
	return createUser, nil
}

func (repo *memkvs) GetByName(name string) (domain.User, error) {
	for _, v := range repo.kvs {
		user := domain.User{}
		err := json.Unmarshal(v, &user)
		if err != nil {
			return domain.User{}, errors.New(apperrors.Internal, err, "fail to get value from kvs", "")
		}
		if user.Name == name {
			return user, nil
		}
	}

	return domain.User{}, nil
}

func (repo *memkvs) Get(id string) (domain.User, error) {
	if value, ok := repo.kvs[id]; ok {
		user := domain.User{}
		err := json.Unmarshal(value, &user)
		if err != nil {
			return domain.User{}, errors.New(apperrors.Internal, err, "fail to get value from kvs", "")
		}

		return user, nil
	}

	return domain.User{}, errors.New(apperrors.NotFound, nil, "user not found in kvs", "")
}

func (repo *memkvs) List() ([]domain.User, error) {
	users := make([]domain.User, 0)
	for _, v := range repo.kvs {
		user := domain.User{}
		err := json.Unmarshal(v, &user)
		if err != nil {
			return []domain.User{}, errors.New(apperrors.Internal, err, "fail to get value from kvs", "")
		}
		users = append(users, user)
	}

	if len(users) > 0 {
		return users, nil
	}

	return []domain.User{}, errors.New(apperrors.NotFound, nil, "user not found in kvs", "")
}

func (repo *memkvs) Update(updateUser domain.User) (domain.User, error) {
	if _, ok := repo.kvs[updateUser.ID]; ok {
		bytes, err := json.Marshal(updateUser)
		if err != nil {
			return domain.User{}, errors.New(apperrors.InvalidInput, err, "user fails at marshal into json string", "")
		}
		repo.kvs[updateUser.ID] = bytes
		return updateUser, nil
	}

	return domain.User{}, errors.New(apperrors.NotFound, nil, "user not found in kvs", "")
}

func (repo *memkvs) Delete(id string) error {
	if _, ok := repo.kvs[id]; ok {
		delete(repo.kvs, id)
		return nil
	}
	return errors.New(apperrors.NotFound, nil, "user not found in kvs", "")
}
