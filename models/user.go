package store

import (
	"github.com/google/uuid"
)

type id uuid.UUID

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"biography"`
	Id        id     `json:"id"`
}

type UserRepo map[id]User

func convertMapToSlice[K comparable, V any](inputMap map[K]V) []V {

}
