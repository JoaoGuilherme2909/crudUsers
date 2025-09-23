package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joaoguilherme2909/crudUsers/store"
	"github.com/joaoguilherme2909/crudUsers/utils"
)

type UserRequest struct {
	Validator utils.Validator `json:"-"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
}

// TODO: adaptar para banco de dados
func NewHandler(db store.UserRepo) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/users", getUsers(db))
	r.Get("/users/{id}", getUserById(db))
	/*	r.Post("/users", addUser(db))
		r.Patch("/users/{id}", update(db))
		r.Delete("/users/{id}", deleteUser(db))
	*/
	return r
}

func getUserById(db store.UserRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		user, err := db.FindById(r.Context(), id)
		if err != nil {
			utils.JsonResponse(w, http.StatusNotFound, map[string]string{
				"errors": "User Not Found",
			})
			return
		}

		utils.JsonResponse(w, http.StatusOK, map[string]store.User{"User": user})
	}
}

func getUsers(db store.UserRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := db.FindAll(r.Context())
		if err != nil {
			utils.JsonResponse(w, http.StatusInternalServerError, map[string]any{"users": users})
			return
		}

		utils.JsonResponse(w, http.StatusOK, map[string]any{"users": users})
	}
}

/*
func addUser(db store.UserRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body UserRequest

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			utils.JsonResponse(w, http.StatusUnprocessableEntity, map[string]any{
				"error": err.Error(),
			})
			return
		}

		body.Validator.CheckField(utils.NotBlank(body.FirstName), "first_name", "This field cannot be blank")
		body.Validator.CheckField(utils.NotBlank(body.LastName), "last_name", "This field cannot be blank")
		body.Validator.CheckField(utils.NotBlank(body.Bio), "bio", "This field cannot be blank")

		body.Validator.CheckField(utils.MinChars(body.FirstName, 2) && utils.MaxChars(body.FirstName, 20), "first_name", "This field must have between 2 and 20 characters")
		body.Validator.CheckField(utils.MinChars(body.LastName, 2) && utils.MaxChars(body.LastName, 20), "last_name", "This field must have between 2 and 20 characters")
		body.Validator.CheckField(utils.MinChars(body.Bio, 20) && utils.MaxChars(body.Bio, 450), "bio", "This field must have between 20 and 450 characters")

		if !body.Validator.Valid() {
			utils.JsonResponse(w, http.StatusBadRequest, body.Validator.FieldErrors)
			return
		}

		user, err := db.Insert(body.FirstName, body.LastName, body.Bio)
		if err != nil {
			utils.JsonResponse(w, http.StatusBadRequest, map[string]any{
				"error": "Could not insert user on database",
			})
			return
		}

		utils.JsonResponse(w, http.StatusCreated, map[string]any{
			"user": user,
		})
	}
}

func update(db store.UserRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var body UserRequest

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			utils.JsonResponse(w, http.StatusUnprocessableEntity, map[string]any{
				"error": err.Error(),
			})
			return
		}

		if utils.NotBlank(body.FirstName) {
			body.Validator.CheckField(utils.MinChars(body.FirstName, 2) && utils.MaxChars(body.FirstName, 20), "first_name", "This field must have between 2 and 20 characters")
		}
		if utils.NotBlank(body.LastName) {
			body.Validator.CheckField(utils.MinChars(body.LastName, 2) && utils.MaxChars(body.LastName, 20), "last_name", "This field must have between 2 and 20 characters")
		}
		if utils.NotBlank(body.Bio) {
			body.Validator.CheckField(utils.MinChars(body.Bio, 20) && utils.MaxChars(body.Bio, 450), "bio", "This field must have between 20 and 450 characters")
		}

		if !body.Validator.Valid() {
			utils.JsonResponse(w, http.StatusBadRequest, body.Validator.FieldErrors)
			return
		}

		user, err := db.Update(id, store.User{
			FirstName: body.FirstName,
			LastName:  body.LastName,
			Bio:       body.Bio,
		})
		if err != nil {
			utils.JsonResponse(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		}

		utils.JsonResponse(w, http.StatusOK, map[string]any{
			"user": user,
		})
	}
}

func deleteUser(db store.UserRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		user, err := db.Delete(id)
		if err != nil {
			utils.JsonResponse(w, http.StatusOK, map[string]any{
				"error": err.Error(),
			})
			return
		}

		utils.JsonResponse(w, http.StatusOK, map[string]any{
			"user": user,
		})
	}
}*/
