package store

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Bio       string    `json:"biography"`
	Id        uuid.UUID `json:"id"`
}

type UserRepo map[uuid.UUID]User

func convertMapToSlice[K comparable, V any](inputMap map[K]V) []V {
	var result []V
	for _, value := range inputMap {
		result = append(result, value)
	}
	return result
}

func (ur UserRepo) FindAll() []User {
	return convertMapToSlice(ur)
}

func (ur UserRepo) FindById(id uuid.UUID) (User, error) {
	user, ok := ur[id]
	if !ok {
		return User{}, errors.New("User not found")
	}

	return user, nil
}
