package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/santigorbe/rest_golang/models"
	"github.com/santigorbe/rest_golang/services"
	"github.com/santigorbe/rest_golang/utils"
)

type TaskHandler struct {
	Service *services.TaskService
}

// Ajusta los handlers según tus funciones reales

// @Summary Obtiene tareas filtradas
// @Description Retorna un listado de tareas filtradas por parámetros
// @Tags tareas
// @Accept  json
// @Produce  json
// @Param limit query integer false "Limit" default(10)
// @Param offset query integer false "Offset" default(0)
// @Param search query string false "Search" default("")
// @Success 200 {object} models.Task
// @Failure 500 {string} string "Error al obtener tareas"
// @Router /tareas [get]
func (h *TaskHandler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Extraer parámetros
	limit, offset := 10, 0
	search := r.URL.Query().Get("search")

	if l, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil {
		limit = l
	}
	if o, err := strconv.Atoi(r.URL.Query().Get("offset")); err == nil {
		offset = o
	}

	// Lógica del servicio
	tasks, total, err := h.Service.GetFilteredTasks(limit, offset, search)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error getting tasks")
		return
	}

	utils.RespondPaginated(w, tasks, total, offset+limit)
}

// @Summary Obtiene un usuario por ID
// @Description Retorna un usuario específico por su ID
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Error getting user"
// @Router /usuarios/{id} [get]
func (h *TaskHandler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	task, err := h.Service.GetTask(id)
	if err != nil {
		if err.Error() == "task not found" {
			utils.RespondError(w, http.StatusNotFound, "Task not found")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, "Error getting task")
		}
		return
	}
	utils.RespondSuccess(w, task)
}

// @Summary Crea un nuevo usuario
// @Description Crea un nuevo usuario con los datos proporcionados
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param user body models.User true "User data"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Error creating user"
// @Router /usuarios [post]
func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error creating task")
		return
	}
	utils.RespondSuccess(w, createdTask)
}

// @Summary Actualiza un usuario existente
// @Description Actualiza un usuario existente con los datos proporcionados
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body models.User true "User data"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Invalid request payload"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Error updating user"
// @Router /usuarios/{id} [put]
func (h *TaskHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	updatedTask, err := h.Service.UpdateTask(id, task)
	if err != nil {
		if err.Error() == "task not found" {
			utils.RespondError(w, http.StatusNotFound, "Task not found")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, "Error updating task")
		}
		return
	}
	utils.RespondSuccess(w, updatedTask)
}

// @Summary Elimina un usuario
// @Description Elimina un usuario por su ID
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {string} string "User deleted successfully"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Error deleting user"
// @Router /usuarios/{id} [delete]
func (h *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := h.Service.DeleteTask(id)
	if err != nil {
		if err.Error() == "task not found" {
			utils.RespondError(w, http.StatusNotFound, "Task not found")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, "Error deleting task")
		}
		return
	}
	utils.RespondSuccess(w, nil)
}
