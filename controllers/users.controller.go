package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/santigorbe/rest_golang/models"
	"github.com/santigorbe/rest_golang/services"
	"github.com/santigorbe/rest_golang/utils"
)

type UserHandler struct {
	Service *services.UserService
}

// @Summary Obtiene todos los usuarios
// @Description Retorna un listado de usuarios
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Failure 500 {string} string "Error getting users"
// @Router /users [get]
func (h *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetUsers()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error getting users")
		return
	}
	utils.RespondSuccess(w, users)
}

// @Summary Obtiene un usuario por ID
// @Description Retorna un usuario específico
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Error getting user"
// @Router /user/{id} [get]
func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := h.Service.GetUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			utils.RespondError(w, http.StatusNotFound, "User not found")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, "Error getting user")
		}
		return
	}
	utils.RespondSuccess(w, user)
}

// @Summary Crea un nuevo usuario
// @Description Crea un nuevo usuario en el sistema
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param user body models.User true "User object"
// @Success 201 {object} models.User
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Error creating user"
// @Router /users [post]
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	createdUser, err := h.Service.CreateUser(user)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error creating user")
		return
	}
	utils.RespondSuccess(w, createdUser)
}

// @Summary Elimina un usuario por ID
// @Description Elimina un usuario específico del sistema
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 204 "No Content"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Error deleting user"
// @Router /user/{id} [delete]
func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := h.Service.DeleteUser(params["id"])
	if err != nil {
		if err.Error() == "user not found" {
			utils.RespondError(w, http.StatusNotFound, "User not found")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, "Error deleting user")
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
	utils.RespondSuccess(w, nil)
}

// @Summary Actualiza un usuario por ID
// @Description Actualiza un usuario específico en el sistema
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body models.User true "User object"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Invalid request payload"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Error updating user"
// @Router /user/{id} [put]
func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	updatedUser, err := h.Service.UpdateUser(params["id"], user)
	if err != nil {
		if err.Error() == "user not found" {
			utils.RespondError(w, http.StatusNotFound, "User not found")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, "Error updating user")
		}
		return
	}
	utils.RespondSuccess(w, updatedUser)
}
