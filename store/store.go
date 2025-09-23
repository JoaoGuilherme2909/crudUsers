package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"biography"`
	Id        string `json:"id"`
}

type UserRepo struct {
	DbConn *pgx.Conn
}

func convertMapToSlice[K comparable, V any](inputMap map[K]V) []V {
	var result []V
	for _, value := range inputMap {
		result = append(result, value)
	}
	return result
}

func (ur UserRepo) FindAll(ctx context.Context) ([]User, error) {
	rows, err := ur.DbConn.Query(ctx, "Select id, first_name, last_name,bio from users")
	if err != nil {
		return []User{}, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Bio,
		)
		if err != nil {
			return []User{}, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return []User{}, fmt.Errorf("Error iterating users rows: %w", err)
	}

	return users, nil
}

func (ur UserRepo) FindById(ctx context.Context, id string) (User, error) {
	sql := `
		SELECT id, first_name, last_name, bio 
		FROM users
		WHERE id = $1
	`

	var user User
	err := ur.DbConn.QueryRow(ctx, sql, id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Bio)
	if err != nil {
		return User{}, fmt.Errorf("Error searching for user: %w", err)
	}

	return user, nil
}

// TODO: Adaptar para a banco de dados
/*
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
}*/
