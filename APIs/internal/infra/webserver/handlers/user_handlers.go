package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/jwtauth"
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/dto"
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/infra/database"
	"github.dev/nicolasmmb/GoExpert-Topicos/pkg/middlewares"
)

type UserHandler struct {
	UserDB database.UserInterface
	JWT    *jwtauth.JWTAuth
	JWTExp int
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: db}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	valid := u.ComparePassword(user.Password)
	if !valid {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	tokenData := map[string]interface{}{
		"id":    u.ID.String(),
		"email": u.Email,
		"name":  u.Name,
	}

	data, err := middlewares.CreateJWTToken(tokenData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output := dto.GetJWTOutput{
		Token: data,
		Type:  "Bearer",
	}
	json.NewEncoder(w).Encode(output)

}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u.ChangePassword(user.Password)
	err = h.UserDB.Create(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userOutput := dto.UserOutputId{
		ID: u.ID.String(),
	}

	json.NewEncoder(w).Encode(userOutput)
	w.WriteHeader(http.StatusCreated)

}

func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userOutput := dto.UserOutput{
		ID:    u.ID.String(),
		Name:  u.Name,
		Email: u.Email,
	}

	json.NewEncoder(w).Encode(userOutput)
	w.WriteHeader(http.StatusOK)

}

func (h *UserHandler) GetByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userOutput := dto.UserOutput{
		ID:    u.ID.String(),
		Name:  u.Name,
		Email: u.Email,
	}

	json.NewEncoder(w).Encode(userOutput)
	w.WriteHeader(http.StatusOK)

}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var user dto.UpdateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err = h.UserDB.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.ChangeName(user.Name)
	u.ChangeEmail(user.Email)
	u.ChangePassword(user.Password)

	err = h.UserDB.Update(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userOutput := dto.UserOutputId{
		ID: u.ID.String(),
	}

	json.NewEncoder(w).Encode(userOutput)
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	err := h.UserDB.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if sort == "" {
		sort = "asc"
	}
	if sort != "asc" && sort != "desc" {
		http.Error(w, "sort must be asc or desc", http.StatusBadRequest)
		return
	}

	users, err := h.UserDB.FindAll(intPage, intLimit, sort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var userOutput []dto.UserOutput

	for _, u := range users {
		userOutput = append(userOutput, dto.UserOutput{
			ID:    u.ID.String(),
			Name:  u.Name,
			Email: u.Email,
		})
	}

	json.NewEncoder(w).Encode(userOutput)
	w.WriteHeader(http.StatusOK)
}
