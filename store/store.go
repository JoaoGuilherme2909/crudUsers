package store

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"biography"`
	Id        string `json:"id"`
}

type UserRepo map[string]User

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

func (ur UserRepo) FindById(id string) (User, error) {
	user, ok := ur[id]
	if !ok {
		return User{}, errors.New("User not found")
	}

	return user, nil
}

func (ur UserRepo) Insert(firstName, lastName, bio string) (User, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		return User{}, err
	}

	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Bio:       bio,
		Id:        id.String(),
	}

	ur[id.String()] = user

	return user, nil
}

func (ur UserRepo) Update(id string, u User) (User, error) {
	user, ok := ur[id]

	if !ok {
		return User{}, errors.New("User not found")
	}

	if u.FirstName != "" {
		user.FirstName = u.FirstName
	}

	if u.LastName != "" {
		user.LastName = u.LastName
	}

	if u.Bio != "" {
		user.Bio = u.Bio
	}

	ur[id] = user

	return user, nil
}

func (ur UserRepo) Delete(id string) (User, error) {
	user, ok := ur[id]

	if !ok {
		return User{}, errors.New("User not found")
	}

	delete(ur, id)

	return user, nil
}
